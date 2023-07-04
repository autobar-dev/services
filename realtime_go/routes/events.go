package routes

import (
	"fmt"

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
		return rest_context.JSON(400, &EventsRouteResponse{
			Status: "error",
			Error:  "session is not valid",
		})
	}

	session_data := utils.ServiceClientTypeToClientType()

	stream_name := rest_context.QueryParam("stream")
	if stream_name == "" {
		new_stream_name := utils.StreamNameFromClientInfo()

		app_context.SseServer.CreateStream()
		fmt.Printf("created a stream for %s\n", session_data.ClientIdentifier)
		return rest_context.Redirect(302, fmt.Sprintf("/events?stream=%s", session_data.ClientIdentifier))
	}

	response_writer := rest_context.Response().Writer
	request := rest_context.Request()

	go func() {
		<-request.Context().Done()
		fmt.Println("client disconnected")
		return
	}()

	app_context.SseServer.ServeHTTP(response_writer, request)

	return nil
}
