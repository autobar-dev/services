package routes

import (
	"github.com/autobar-dev/services/module/controllers"
	"github.com/autobar-dev/services/module/types"
	"github.com/labstack/echo/v4"
)

type ActivateRouteRequestBody struct {
	SerialNumber string `json:"serial_number"`
	Otk          string `json:"otk"` // One Time Key
}

type ActivateRouteResponse struct {
	Status string  `json:"status"`
	Error  *string `json:"error"`
}

func ActivateRoute(c echo.Context) error {
	rest_context := c.(*types.RestContext)
	app_context := rest_context.AppContext
	client_context := rest_context.ClientContext

	if client_context == nil {
		err := "not authorized"
		return rest_context.JSON(401, &ActivateRouteResponse{
			Status: "error",
			Error:  &err,
		})
	}

	var arrb ActivateRouteRequestBody
	if err := c.Bind(&arrb); err != nil {
		err := err.Error()
		return rest_context.JSON(401, &ActivateRouteResponse{
			Status: "error",
			Error:  &err,
		})
	}

	err := controllers.ActivateController(app_context, client_context.Identifier, arrb.SerialNumber, arrb.Otk)
	if err != nil {
		err := err.Error()
		return rest_context.JSON(400, &ActivateRouteResponse{
			Status: "error",
			Error:  &err,
		})
	}

	return rest_context.JSON(200, &CreateModuleRouteResponse{
		Status: "ok",
		Error:  nil,
	})
}
