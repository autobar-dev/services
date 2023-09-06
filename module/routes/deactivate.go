package routes

import (
	"github.com/autobar-dev/services/module/controllers"
	"github.com/autobar-dev/services/module/types"
	"github.com/labstack/echo/v4"
)

type DeactivateRouteResponse struct {
	Status string  `json:"status"`
	Error  *string `json:"error"`
}

func DeactivateRoute(c echo.Context) error {
	rest_context := c.(*types.RestContext)
	app_context := rest_context.AppContext
	client_context := rest_context.ClientContext

	if client_context == nil {
		err := "not authorized"
		return rest_context.JSON(401, &DeactivateRouteResponse{
			Status: "error",
			Error:  &err,
		})
	}

	err := controllers.DeactivateController(app_context, client_context.Identifier)
	if err != nil {
		err := err.Error()
		return rest_context.JSON(400, &DeactivateRouteResponse{
			Status: "error",
			Error:  &err,
		})
	}

	return rest_context.JSON(200, &DeactivateRouteResponse{
		Status: "ok",
		Error:  nil,
	})
}
