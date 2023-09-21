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

	ct := utils.ServiceClientTypeToClientType(client_context.Type)
	id := client_context.Identifier

	exchange_name := utils.ExchangeNameFromClientInfo(ct, id)

	queue_name, err := mr.CreatePubSub(exchange_name)
	if err != nil {
		return rest_context.JSON(500, &WsRouteResponse{
			Status: "error",
			Error:  "failed to create pubsub",
		})
	}

	fmt.Printf("will try to consume %s\n", *queue_name)

	queue_consumer_name := utils.QueueConsumerName(*queue_name)
	commands, err := mr.ConsumeCommands(*queue_name, "") // consumer name auto-generated
	if err != nil {
		_ = mr.CancelConsumer(queue_consumer_name)
		return rest_context.JSON(500, &WsRouteResponse{
			Status: "error",
			Error:  "failed to consume queue",
		})
	}

	err = sr.IncrementListenersCountForExchange(exchange_name)
	if err != nil {
		_ = mr.CancelConsumer(queue_consumer_name)
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
	defer heartbeat_ticker.Stop()

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

		close(received)
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

		close(to_send)
	}()

	for {
		_, message, err := ws.ReadMessage()
		if err != nil {
			fmt.Printf("failed to read message: %s\n", err)
			break
		}

		received <- string(message)
	}

	_ = mr.CancelConsumer(queue_consumer_name)
	_ = sr.DecrementListenersCountForExchange(exchange_name)

	return nil
}
