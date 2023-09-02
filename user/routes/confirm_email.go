package routes

import (
	"github.com/autobar-dev/services/user/controllers"
	"github.com/autobar-dev/services/user/types"
	"github.com/autobar-dev/services/user/utils"
	"github.com/labstack/echo/v4"
)

type ConfirmEmailRequestBody struct {
	ConfirmationCode string `json:"confirmation_code"`
	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name"`
	DateOfBirth      string `json:"date_of_birth"`
	CurrencyCode     string `json:"currency_code"`
}

type ConfirmEmailRouteResponse struct {
	Status string  `json:"status"`
	Error  *string `json:"error"`
}

func ConfirmEmailRoute(c echo.Context) error {
	rest_context := c.(*types.RestContext)
	app_context := *(*rest_context).AppContext

	var cerb ConfirmEmailRequestBody
	err := c.Bind(&cerb)
	if err != nil {
		err := err.Error()
		return c.JSON(400, &CreateUserRouteResponse{
			Status: "error",
			Error:  &err,
		})
	}

	dob, err := utils.DateStringToTime(cerb.DateOfBirth)
	if err != nil {
		err := err.Error()
		return c.JSON(400, &CreateUserRouteResponse{
			Status: "error",
			Error:  &err,
		})
	}

	err = controllers.ConfirmEmail(
		&app_context,
		cerb.ConfirmationCode,
		cerb.FirstName,
		cerb.LastName,
		dob,
		cerb.CurrencyCode,
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
