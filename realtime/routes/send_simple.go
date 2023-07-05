package routes

import (
	"github.com/labstack/echo/v4"

	"go.a5r.dev/services/realtime/controllers"
	"go.a5r.dev/services/realtime/types"
	"go.a5r.dev/services/realtime/utils"
)

type SendSimpleRouteBody struct {
	ClientType string `json:"client_type"`
	Identifier string `json:"identifier"`
	Body       string `json:"body"`
}

type SendSimpleRouteResponse struct {
	Status string  `json:"status"`
	Error  *string `json:"error"`
}

func SendSimpleRoute(c echo.Context) error {
	rest_context := c.(*types.RestContext)
	app_context := *(*rest_context).AppContext

	var ssrb SendSimpleRouteBody
	err := rest_context.Bind(&ssrb)
	if err != nil {
		return rest_context.JSON(400, &SendSimpleRouteResponse{
			Status: "error",
			Error:  nil,
		})
	}

	var ct types.ClientType

	ctp, err := utils.ClientTypeFromString(ssrb.ClientType)
	if err != nil {
		err := "client_type invalid"
		return rest_context.JSON(400, &SendSimpleRouteResponse{
			Status: "error",
			Error:  &err,
		})
	}

	ct = *ctp

	err = controllers.SendSimpleMessage(app_context, ct, ssrb.Identifier, ssrb.Body)
	if err != nil {
		err := err.Error()

		return rest_context.JSON(400, &SendSimpleRouteResponse{
			Status: "error",
			Error:  &err,
		})
	}

	return rest_context.JSON(200, &SendSimpleRouteResponse{
		Status: "ok",
		Error:  nil,
	})
}
