package routes

import (
	"github.com/autobar-dev/services/user/controllers"
	"github.com/autobar-dev/services/user/types"
	"github.com/labstack/echo/v4"
)

type WhoAmIRouteResponse struct {
	Status string      `json:"status"`
	Error  *string     `json:"error"`
	Data   *types.User `json:"data"`
}

func WhoAmIRoute(c echo.Context) error {
	rest_context := c.(*types.RestContext)
	app_context := *(*rest_context).AppContext
	client_context := rest_context.ClientContext

	if client_context == nil {
		err := "not authenticated"
		return c.JSON(401, &WhoAmIRouteResponse{
			Status: "error",
			Error:  &err,
			Data:   nil,
		})
	}

	user, err := controllers.GetUserById(&app_context, client_context.Identifier)
	if err != nil {
		err := err.Error()
		return c.JSON(400, &WhoAmIRouteResponse{
			Status: "error",
			Error:  &err,
			Data:   nil,
		})
	}

	return c.JSON(200, &WhoAmIRouteResponse{
		Status: "ok",
		Error:  nil,
		Data:   user,
	})
}
