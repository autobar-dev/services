package routes

import (
	"github.com/autobar-dev/services/auth/controllers"
	"github.com/autobar-dev/services/auth/types"
	"github.com/labstack/echo/v4"
)

type RefreshRequestBody struct {
	RefreshToken string `json:"refresh_token"`
}

type RefreshResponse struct {
	Status string        `json:"status"`
	Error  *string       `json:"error"`
	Data   *types.Tokens `json:"data"`
}

func RefreshRoute(c echo.Context) error {
	rest_context := c.(*types.RestContext)
	app_context := *(*rest_context).AppContext

	var rrb RefreshRequestBody
	err := c.Bind(&rrb)
	if err != nil {
		err := "incorrect body format"
		return rest_context.JSON(400, &RefreshResponse{
			Status: "error",
			Error:  &err,
			Data:   nil,
		})
	}

	if rrb.RefreshToken == "" {
		err := "missing refresh token"
		return rest_context.JSON(400, &RefreshResponse{
			Status: "error",
			Error:  &err,
			Data:   nil,
		})
	}

	tokens, err := controllers.Refresh(&app_context, rrb.RefreshToken)
	if err != nil {
		err := "failed to refresh tokens"
		return rest_context.JSON(400, &RefreshResponse{
			Status: "error",
			Error:  &err,
			Data:   nil,
		})
	}

	return rest_context.JSON(200, &RefreshResponse{
		Status: "ok",
		Error:  nil,
		Data:   tokens,
	})
}
