package routes

import (
	"github.com/labstack/echo/v4"
	"go.a5r.dev/services/module/controllers"
	"go.a5r.dev/services/module/types"
)

type GetModuleRouteResponse struct {
	Status string        `json:"status"`
	Data   *types.Module `json:"data"`
	Error  *string       `json:"error"`
}

func GetModuleRoute(c echo.Context) error {
	rest_context := c.(*types.RestContext)
	app_context := *(*rest_context).AppContext

	serial_number := rest_context.QueryParam("serial_number")
	if serial_number == "" {
		err := "missing serial_number query argument"

		return rest_context.JSON(400, &GetModuleRouteResponse{
			Status: "error",
			Data:   nil,
			Error:  &err,
		})
	}

	module, err := controllers.GetModuleController(&app_context, serial_number)
	if err != nil {
		err := err.Error()

		return rest_context.JSON(400, &GetModuleRouteResponse{
			Status: "error",
			Data:   nil,
			Error:  &err,
		})
	}

	return rest_context.JSON(200, &GetModuleRouteResponse{
		Status: "ok",
		Data:   module,
		Error:  nil,
	})
}
