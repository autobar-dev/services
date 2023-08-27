package routes

import (
	"fmt"

	"github.com/labstack/echo/v4"

	"github.com/autobar-dev/services/realtime/controllers"
	"github.com/autobar-dev/services/realtime/types"
	"github.com/autobar-dev/services/realtime/utils"
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
	client_context := rest_context.ClientContext

	if client_context == nil {
		err := "unauthorized"
		return rest_context.JSON(400, &ReplyRouteResponse{
			Status: "error",
			Error:  &err,
		})
	}

	var rrb ReplyRouteBody
	err := rest_context.Bind(&rrb)
	if err != nil {
		return rest_context.JSON(400, &ReplyRouteResponse{
			Status: "error",
			Error:  nil,
		})
	}

	fmt.Println("pre scttct")

	ct := utils.ServiceClientTypeToClientType(client_context.Type)

	fmt.Printf("will try to reply to #%s as %s (%s)\n", rrb.Id, client_context.Identifier, ct)

	err = controllers.Reply(app_context, ct, client_context.Identifier, rrb.Id)
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
