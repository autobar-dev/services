package routes

import (
	"github.com/autobar-dev/services/auth/types"
	"github.com/labstack/echo/v4"
)

func IsValidRoute(c echo.Context) error {
	rest_context := c.(*types.RestContext)
	client_context := rest_context.ClientContext

	if client_context == nil {
		return c.JSON(401, map[string]interface{}{
			"error": "unauthorized",
		})
	}

	return c.JSON(200, client_context)
}
