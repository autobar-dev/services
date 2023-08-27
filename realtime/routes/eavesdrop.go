package routes

import (
	"fmt"

	"github.com/autobar-dev/services/realtime/types"
	"github.com/autobar-dev/services/realtime/utils"
	"github.com/labstack/echo/v4"
)

type EavesdropRouteResponse struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

func EavesdropRoute(c echo.Context) error {
	rest_context := c.(*types.RestContext)
	app_context := *(*rest_context).AppContext

	rr := app_context.Repositories.Redis

	identifier := rest_context.QueryParam("identifier")
	client_type := rest_context.QueryParam("client_type")

	var ct types.ClientType

	if identifier == "" || client_type == "" {
		return rest_context.JSON(400, &EavesdropRouteResponse{
			Status: "error",
			Error:  "identifier or client_type not provided",
		})
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

	exchange_name := utils.ExchangeNameFromClientInfo(ct, identifier)

	listeners, err := rr.GetListenersCountForExchange(exchange_name)
	if err != nil {
		return rest_context.JSON(500, &EavesdropRouteResponse{
			Status: "error",
			Error:  "failed to retrieve listeners",
		})
	}

	if *listeners == 0 {
		return rest_context.JSON(400, &EavesdropRouteResponse{
			Status: "error",
			Error:  "no client listeners connected",
		})
	}

	redirect_uri := fmt.Sprintf(
		"%s/events?identifier=%s&client_type=%s",
		app_context.Config.ServiceBasepath,
		identifier,
		client_type,
	)

	return rest_context.Redirect(302, redirect_uri)
}
