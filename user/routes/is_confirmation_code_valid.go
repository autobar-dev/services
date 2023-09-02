package routes

import (
	"github.com/autobar-dev/services/user/controllers"
	"github.com/autobar-dev/services/user/types"
	"github.com/labstack/echo/v4"
)

type IsConfirmationCodeValidRequestQuery struct {
	ConfirmationCode string `query:"code"`
}

func IsConfirmationCodeValidRoute(c echo.Context) error {
	rest_context := c.(*types.RestContext)
	app_context := *(*rest_context).AppContext

	var iccvrq IsConfirmationCodeValidRequestQuery
	err := c.Bind(&iccvrq)
	if err != nil {
		err := err.Error()
		return c.JSON(400, &CreateUserRouteResponse{
			Status: "error",
			Error:  &err,
		})
	}

	err = controllers.IsConfirmationCodeValid(
		&app_context,
		iccvrq.ConfirmationCode,
	)
	if err != nil {
		err := err.Error()
		return c.JSON(400, &CreateUserRouteResponse{
			Status: "error",
			Error:  &err,
		})
	}

	return c.JSON(200, &CreateUserRouteResponse{
		Status: "ok",
		Error:  nil,
	})
}
