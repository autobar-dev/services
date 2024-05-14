package routes

import (
	"fmt"

	"github.com/autobar-dev/services/file/controllers"
	"github.com/autobar-dev/services/file/types"
	"github.com/labstack/echo/v4"
)

type DeleteRouteResponse struct {
	Status string  `json:"status"`
	Error  *string `json:"error"`
}

type DeleteRouteRequestBody struct {
	Id string `json:"id"`
}

func DeleteRoute(c echo.Context) error {
	rest_context := c.(*types.RestContext)
	app_context := *(*rest_context).AppContext

	var drrb DeleteRouteRequestBody
	err := rest_context.Bind(&drrb)
	if err != nil {
		err := "failed to parse request body"
		return rest_context.JSON(400, &DeleteRouteResponse{
			Status: "error",
			Error:  &err,
		})
	}

	err = controllers.DeleteFile(&app_context, drrb.Id)
	if err != nil {
		fmt.Println(err)
		err := "failed to delete file"
		return rest_context.JSON(500, &DeleteRouteResponse{
			Status: "error",
			Error:  &err,
		})
	}

	return rest_context.JSON(200, &DeleteRouteResponse{
		Status: "ok",
		Error:  nil,
	})
}
