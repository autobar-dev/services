package routes

import (
	"github.com/labstack/echo/v4"
	"go.a5r.dev/services/module/controllers"
	"go.a5r.dev/services/module/types"
)

type GetAllModulesRouteResponse struct {
	Status string          `json:"status"`
	Data   *[]types.Module `json:"data"`
	Error  *string         `json:"error"`
}

func GetAllModulesRoute(c echo.Context) error {
	rest_context := c.(*types.RestContext)
	app_context := *(*rest_context).AppContext

	modules, err := controllers.GetAllModulesController(&app_context)
	if err != nil {
		err := err.Error()

		return rest_context.JSON(400, &GetAllModulesRouteResponse{
			Status: "error",
			Data:   nil,
			Error:  &err,
		})
	}

	return rest_context.JSON(200, &GetAllModulesRouteResponse{
		Status: "ok",
		Data:   modules,
		Error:  nil,
	})
}
