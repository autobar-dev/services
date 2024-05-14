package routes

import (
	"github.com/autobar-dev/services/file/controllers"
	"github.com/autobar-dev/services/file/types"
	"github.com/labstack/echo/v4"
)

type UploadRouteResponseData struct {
	Id string `json:"id"`
}

type UploadRouteResponse struct {
	Status string                   `json:"status"`
	Error  *string                  `json:"error"`
	Data   *UploadRouteResponseData `json:"data"`
}

func UploadRoute(c echo.Context) error {
	rest_context := c.(*types.RestContext)
	app_context := *(*rest_context).AppContext

	file, err := c.FormFile("file")
	if err != nil {
		err := err.Error()
		return c.JSON(500, &UploadRouteResponse{
			Status: "error",
			Error:  &err,
			Data:   nil,
		})
	}

	id, err := controllers.UploadFile(&app_context, file)
	if err != nil {
		err := err.Error()
		return c.JSON(500, &UploadRouteResponse{
			Status: "error",
			Error:  &err,
			Data:   nil,
		})
	}

	return c.JSON(200, &UploadRouteResponse{
		Status: "ok",
		Error:  nil,
		Data: &UploadRouteResponseData{
			Id: id,
		},
	})
}
