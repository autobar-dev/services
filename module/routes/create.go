package routes

import (
	"github.com/autobar-dev/services/module/controllers"
	"github.com/autobar-dev/services/module/types"
	"github.com/labstack/echo/v4"
)

type CreateModuleRouteResponse struct {
	Status string                      `json:"status"`
	Data   *types.CreateModuleResponse `json:"data"`
	Error  *string                     `json:"error"`
}

func CreateModuleRoute(c echo.Context) error {
	rest_context := c.(*types.RestContext)
	app_context := *(*rest_context).AppContext

	module, err := controllers.CreateModuleController(&app_context)
	if err != nil {
		err := err.Error()

		return rest_context.JSON(400, &CreateModuleRouteResponse{
			Status: "error",
			Data:   nil,
			Error:  &err,
		})
	}

	return rest_context.JSON(200, &CreateModuleRouteResponse{
		Status: "ok",
		Data:   module,
		Error:  nil,
	})
}
