package routes

import (
	"github.com/autobar-dev/services/user/controllers"
	"github.com/autobar-dev/services/user/types"
	"github.com/labstack/echo/v4"
)

type CreateUserRequestBody struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	DateOfBirth string `json:"date_of_birth"`
	Locale      string `json:"locale"`
}

type CreateUserRouteResponse struct {
	Status string  `json:"status"`
	Error  *string `json:"error"`
}

func CreateUserRoute(c echo.Context) error {
	rest_context := c.(*types.RestContext)
	app_context := *(*rest_context).AppContext

	var curb CreateUserRequestBody
	err := c.Bind(&curb)
	if err != nil {
		err := err.Error()
		return c.JSON(400, &CreateUserRouteResponse{
			Status: "error",
			Error:  &err,
		})
	}

	err = controllers.Register(
		&app_context,
		curb.Email,
		curb.Password,
		curb.FirstName,
		curb.LastName,
		curb.DateOfBirth,
		curb.Locale,
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
