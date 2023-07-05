package routes

import (
	"github.com/labstack/echo/v4"

	"go.a5r.dev/services/realtime/controllers"
	"go.a5r.dev/services/realtime/types"
	"go.a5r.dev/services/realtime/utils"
)

type SendCommandRouteBody struct {
	ClientType string `json:"client_type"`
	Identifier string `json:"identifier"`
	Command    string `json:"command"`
	Args       string `json:"args"`
}

type SendCommandRouteResponse struct {
	Status string  `json:"status"`
	Error  *string `json:"error"`
}

func SendCommandRoute(c echo.Context) error {
	rest_context := c.(*types.RestContext)
	app_context := *(*rest_context).AppContext

	var scrb SendCommandRouteBody
	err := rest_context.Bind(&scrb)
	if err != nil {
		return rest_context.JSON(400, &SendCommandRouteResponse{
			Status: "error",
			Error:  nil,
		})
	}

	var ct types.ClientType

	ctp, err := utils.ClientTypeFromString(scrb.ClientType)
	if err != nil {
		err := "client_type invalid"
		return rest_context.JSON(400, &SendCommandRouteResponse{
			Status: "error",
			Error:  &err,
		})
	}

	ct = *ctp

	err = controllers.SendCommandMessage(app_context, ct, scrb.Identifier, scrb.Command, scrb.Args)
	if err != nil {
		err := err.Error()

		return rest_context.JSON(400, &SendCommandRouteResponse{
			Status: "error",
			Error:  &err,
		})
	}

	return rest_context.JSON(200, &SendCommandRouteResponse{
		Status: "ok",
		Error:  nil,
	})
}
