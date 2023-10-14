package routes

import (
	"fmt"
	"time"

	"github.com/autobar-dev/services/realtime/types"
	"github.com/autobar-dev/services/realtime/utils"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

type WsRouteResponse struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

var upgrader = websocket.Upgrader{}

func WsRoute(c echo.Context) error {
	rest_context := c.(*types.RestContext)
	app_context := rest_context.AppContext
	client_context := rest_context.ClientContext

	if client_context == nil {
		return rest_context.JSON(401, &WsRouteResponse{
			Status: "error",
			Error:  "unauthorized",
		})
	}

	mr := app_context.Repositories.Mq
	sr := app_context.Repositories.State

	amqp_channel, err := app_context.AmqpConnection.Channel()
	if err != nil {
		panic(fmt.Sprintf("failed to open a channel: %s", err))
	}
	defer amqp_channel.Close()

	ct := utils.ServiceClientTypeToClientType(client_context.Type)
	id := client_context.Identifier

	exchange_name := utils.ExchangeNameFromClientInfo(ct, id)

	queue_name, err := mr.CreatePubSub(amqp_channel, exchange_name)
	if err != nil {
		return rest_context.JSON(500, &WsRouteResponse{
			Status: "error",
			Error:  "failed to create pubsub",
		})
	}

	queue_consumer_name := utils.QueueConsumerName(*queue_name)
	commands, err := mr.ConsumeCommands(amqp_channel, *queue_name, queue_consumer_name) // consumer name auto-generated
	if err != nil {
		_ = mr.CancelConsumer(amqp_channel, queue_consumer_name)
		return rest_context.JSON(500, &WsRouteResponse{
			Status: "error",
			Error:  "failed to consume queue",
		})
	}

	fmt.Printf("consuming queue %s for %s\n", *queue_name, id)

	err = sr.IncrementListenersCountForExchange(exchange_name)
	if err != nil {
		_ = mr.CancelConsumer(amqp_channel, queue_consumer_name)
		return rest_context.JSON(500, &WsRouteResponse{
			Status: "error",
			Error:  "failed to increment listeners count",
		})
	}

	// upgrade to websocket
	ws, err := upgrader.Upgrade(c.Response().Writer, c.Request(), nil)
	if err != nil {
		fmt.Printf("failed to upgrade: %s\n", err)
		return err
	}
	defer ws.Close()

	heartbeat_ticker := time.NewTicker(10 * time.Second)

	received := make(chan string)
	to_send := make(chan *types.Command)

	// heartbeat goroutine
	go func() {
		for range heartbeat_ticker.C {
			to_send <- &types.Command{
				Id:      "",
				Command: "heartbeat",
				Args: map[string]interface{}{
					"timestamp": time.Now().Unix(),
				},
			}
		}
	}()

	// from message queue goroutine
	go func() {
		for delivery := range commands {
			to_send <- &types.Command{
				Id:      delivery.Id,
				Command: delivery.Command,
				Args:    delivery.Args,
			}
		}
	}()

	// receiver goroutine
	go func() {
		for received_message := range received {
			fmt.Printf("received message: %s\n", received_message)
		}

		// close(received)
	}()

	// sender goroutine
	go func() {
		for message_to_send := range to_send {
			err := ws.WriteJSON(message_to_send)
			if err != nil {
				fmt.Printf("failed to send message: %s\n", err)
				break
			}
		}

		// close(to_send)
	}()

	for {
		_, message, err := ws.ReadMessage()
		if err != nil {
			fmt.Printf("failed to read message: %s\n", err)
			break
		}

		received <- string(message)
	}

	heartbeat_ticker.Stop()
	close(received)
	close(to_send)

	err = mr.CancelConsumer(amqp_channel, queue_consumer_name)
	if err != nil {
		fmt.Printf("WARNING: MQ err: %s\n", err)
	}
	_ = sr.DecrementListenersCountForExchange(exchange_name)

	return nil
}
