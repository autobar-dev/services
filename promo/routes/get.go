package routes

import (
	"github.com/autobar-dev/services/file/controllers"
	"github.com/autobar-dev/services/file/types"
	"github.com/labstack/echo/v4"
)

type GetFileRouteResponse struct {
	Status string      `json:"status"`
	Error  *string     `json:"error"`
	Data   *types.File `json:"data"`
}

func GetFileRoute(c echo.Context) error {
	rest_context := c.(*types.RestContext)
	app_context := *(*rest_context).AppContext

	id := c.QueryParam("id")
	download := c.QueryParam("download") == "1"

	if id == "" {
		err := "id must be provided"
		return rest_context.JSON(400, &GetFileRouteResponse{
			Status: "error",
			Error:  &err,
			Data:   nil,
		})
	}

	file, err := controllers.GetFile(&app_context, id, download)
	if err != nil {
		err := err.Error()
		return c.JSON(400, &GetFileRouteResponse{
			Status: "error",
			Error:  &err,
			Data:   nil,
		})
	}

	return c.JSON(200, &GetFileRouteResponse{
		Status: "ok",
		Error:  nil,
		Data:   file,
	})
}
