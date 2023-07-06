package routes

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go.a5r.dev/services/realtime/types"
	"go.a5r.dev/services/realtime/utils"
)

type EventsRouteResponse struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

func EventsRoute(c echo.Context) error {
	rest_context := c.(*types.RestContext)
	app_context := *(*rest_context).AppContext

	ar := app_context.Repositories.Auth
	mr := app_context.Repositories.Mq
	rr := app_context.Repositories.Redis

	session_from_query := rest_context.QueryParam("session")
	session_from_header := rest_context.Request().Header.Get("session")
	session_from_cookie, _ := rest_context.Request().Cookie("session_id")
	session := ""

	if session_from_query != "" {
		session = session_from_query
	} else if session_from_header != "" {
		session = session_from_header
	} else if session_from_cookie != nil {
		session = session_from_cookie.Value
	} else {
		return rest_context.JSON(401, &EventsRouteResponse{
			Status: "error",
			Error:  "session not provided",
		})
	}

	session_data, err := ar.VerifySession(session)
	if err != nil {
		fmt.Printf("error while verifying session: %+v\n", err)

		return rest_context.JSON(400, &EventsRouteResponse{
			Status: "error",
			Error:  "session is not valid",
		})
	}

	stream_name := rest_context.QueryParam("stream")
	identifier := rest_context.QueryParam("identifier")
	client_type := rest_context.QueryParam("client_type")

	var ct types.ClientType

	if client_type == "" {
		ct = utils.ServiceClientTypeToClientType(session_data.ClientType)
	} else {
		ctp, err := utils.ClientTypeFromString(client_type)
		if err != nil {
			return rest_context.JSON(400, &EventsRouteResponse{
				Status: "error",
				Error:  "client_type not provided",
			})
		}

		ct = *ctp
	}

	var id string

	if identifier == "" {
		id = session_data.ClientIdentifier
	} else {
		id = identifier
	}

	if stream_name == "" {
		unique_id := uuid.New().String()
		new_stream_name := utils.ExchangeNameFromClientInfo(ct, id) + fmt.Sprintf("_%s", unique_id)

		app_context.SseServer.CreateStream(new_stream_name)

		return rest_context.Redirect(302, fmt.Sprintf("/events?stream=%s&identifier=%s&client_type=%s&session=%s", new_stream_name, identifier, client_type, session_from_query))
	}

	// There is a listener at this point
	exchange_name := utils.ExchangeNameFromClientInfo(ct, id)

	queue_name, err := mr.CreatePubSubQueue(exchange_name)
	if err != nil {
		return rest_context.JSON(500, &EventsRouteResponse{
			Status: "error",
			Error:  "failed to establish queue",
		})
	}

	fmt.Printf("will try to consume %s\n", *queue_name)

	deliveries, err := mr.Consume(*queue_name)
	if err != nil {
		return rest_context.JSON(500, &EventsRouteResponse{
			Status: "error",
			Error:  "failed to consume queue",
		})
	}

	err = rr.IncrementListenersCountForExchange(exchange_name)
	if err != nil {
		return rest_context.JSON(500, &EventsRouteResponse{
			Status: "error",
			Error:  "failed to register listener in Redis",
		})
	}

	response_writer := rest_context.Response().Writer
	request := rest_context.Request()

	heartbeat_interval, _ := time.ParseDuration(fmt.Sprintf("%ds", app_context.Config.SseHeartbeatIntervalSeconds))
	heartbeart_ticker := time.NewTicker(heartbeat_interval)
	heartbeat_done := make(chan bool)

	listen_done := make(chan bool)

	go func() {
		for {
			select {
			case <-heartbeat_done:
				heartbeart_ticker.Stop()
				return
			case <-heartbeart_ticker.C:
				app_context.SseServer.Publish(stream_name, utils.CreateHeartbeatSseEvent())
			}
		}
	}()

	go func() {
		for {
			select {
			case <-listen_done:
				return
			case delivery := <-deliveries:
				var message map[string]string

				err := json.Unmarshal(delivery.Body, &message)
				if err != nil {
					fmt.Printf("failed to parse delivery from queue: %+v\n", err)
				} else {

				}

				message_type := message["type"]

				switch message_type {
				case "simple":
					simple_message_body := message["body"]

					app_context.SseServer.Publish(stream_name, utils.CreateSimpleSseEvent(simple_message_body))
					break
				case "command":
					command_message_command := message["command"]
					command_message_args := message["args"]

					app_context.SseServer.Publish(stream_name, utils.CreateCommandSseEvent(command_message_command, command_message_args))
					break
				default:
					fmt.Printf("unknown message type from queue: '%s'\n", message["type"])
				}
			}
		}
	}()

	go func() {
		<-request.Context().Done()

		heartbeat_done <- true
		listen_done <- true

		err = rr.DecrementListenersCountForExchange(exchange_name)
		if err != nil {
			fmt.Printf("error while decrementing listeners count (key=%s): %+v\n", exchange_name, err)
		}

		fmt.Println("client disconnected")
		return
	}()

	app_context.SseServer.ServeHTTP(response_writer, request)

	return nil
}
