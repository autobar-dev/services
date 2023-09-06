package routes

import (
	"github.com/autobar-dev/services/user/controllers"
	"github.com/autobar-dev/services/user/types"
	"github.com/labstack/echo/v4"
)

type GetUserRouteResponse struct {
	Status string      `json:"status"`
	Error  *string     `json:"error"`
	Data   *types.User `json:"data"`
}

func GetUserRoute(c echo.Context) error {
	rest_context := c.(*types.RestContext)
	app_context := *(*rest_context).AppContext

	id := c.QueryParam("id")

	// Id not specified
	if id == "" {
		err := "id must be provided"
		return c.JSON(400, &GetUserRouteResponse{
			Status: "error",
			Error:  &err,
			Data:   nil,
		})
	}

	user, err := controllers.GetUserById(&app_context, id)
	if err != nil {
		err := err.Error()
		return c.JSON(400, &GetUserRouteResponse{
			Status: "error",
			Error:  &err,
			Data:   nil,
		})
	}

	return c.JSON(200, &GetUserRouteResponse{
		Status: "ok",
		Error:  nil,
		Data:   user,
	})
}
