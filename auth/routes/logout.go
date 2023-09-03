package routes

import (
	"github.com/autobar-dev/services/auth/controllers"
	"github.com/autobar-dev/services/auth/types"
	"github.com/labstack/echo/v4"
)

type LogoutRequestBody struct {
	RefreshToken string `json:"refresh_token"`
}

type LogoutResponse struct {
	Status string  `json:"status"`
	Error  *string `json:"error"`
}

func LogoutRoute(c echo.Context) error {
	rest_context := c.(*types.RestContext)
	app_context := *(*rest_context).AppContext

	var lrb LogoutRequestBody
	err := c.Bind(&lrb)
	if err != nil {
		err := "incorrect body format"
		return rest_context.JSON(400, &LogoutResponse{
			Status: "error",
			Error:  &err,
		})
	}

	err = controllers.Logout(&app_context, lrb.RefreshToken)
	if err != nil {
		err := "failed to logout"
		return rest_context.JSON(400, &LogoutResponse{
			Status: "error",
			Error:  &err,
		})
	}

	return rest_context.JSON(200, &LogoutResponse{
		Status: "ok",
		Error:  nil,
	})
}
