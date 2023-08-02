package routes

import (
	"fmt"

	"github.com/labstack/echo/v4"

	"go.a5r.dev/services/realtime/controllers"
	"go.a5r.dev/services/realtime/types"
	"go.a5r.dev/services/realtime/utils"
)

type ReplyRouteBody struct {
	Id string `json:"id"`
}

type ReplyRouteResponse struct {
	Status string  `json:"status"`
	Error  *string `json:"error"`
}

func ReplyRoute(c echo.Context) error {
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
		err := "session not provided"
		return rest_context.JSON(401, &ReplyRouteResponse{
			Status: "error",
			Error:  &err,
		})
	}

	fmt.Println("theres been a request to reply")

	session_data, err := ar.VerifySession(session)
	if err != nil {
		fmt.Printf("error while verifying session: %+v\n", err)

		err := "session is not valid"
		return rest_context.JSON(400, &ReplyRouteResponse{
			Status: "error",
			Error:  &err,
		})
	}

	fmt.Printf("session data in request to reply: %v\n", session_data)

	var rrb ReplyRouteBody
	err = rest_context.Bind(&rrb)
	if err != nil {
		return rest_context.JSON(400, &ReplyRouteResponse{
			Status: "error",
			Error:  nil,
		})
	}

	fmt.Println("pre scttct")

	ct := utils.ServiceClientTypeToClientType(session_data.ClientType)

	fmt.Printf("will try to reply to #%s as %s (%s)\n", rrb.Id, session_data.ClientIdentifier, ct)

	err = controllers.Reply(app_context, ct, session_data.ClientIdentifier, rrb.Id)
	if err != nil {
		err := err.Error()

		return rest_context.JSON(400, &ReplyRouteResponse{
			Status: "error",
			Error:  &err,
		})
	}

	return rest_context.JSON(200, &ReplyRouteResponse{
		Status: "ok",
		Error:  nil,
	})
}
