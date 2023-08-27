package routes

import (
	"github.com/autobar-dev/services/module/controllers"
	"github.com/autobar-dev/services/module/types"
	"github.com/labstack/echo/v4"
)

type GetAllModulesForStationRouteResponse struct {
	Status string          `json:"status"`
	Data   *[]types.Module `json:"data"`
	Error  *string         `json:"error"`
}

func GetAllModulesForStationRoute(c echo.Context) error {
	rest_context := c.(*types.RestContext)
	app_context := *(*rest_context).AppContext

	station_id := rest_context.QueryParam("station_id")
	if station_id == "" {
		err := "station_id needs to be provided as a query parameter"
		return rest_context.JSON(400, &GetAllModulesForStationRouteResponse{
			Status: "error",
			Data:   nil,
			Error:  &err,
		})
	}

	modules, err := controllers.GetAllModulesForStationController(&app_context, station_id)
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
