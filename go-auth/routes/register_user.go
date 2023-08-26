package routes

import (
	"github.com/autobar-dev/services/auth/controllers"
	"github.com/autobar-dev/services/auth/types"
	"github.com/labstack/echo/v4"
)

type RegisterUserRequestBody struct {
	UserId   string `json:"user_id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterUserResponse struct {
	Status string  `json:"status"`
	Error  *string `json:"error"`
	Data   *string `json:"data"`
}

func RegisterUserRoute(c echo.Context) error {
	rest_context := c.(*types.RestContext)
	app_context := *(*rest_context).AppContext

	var rurb RegisterUserRequestBody
	err := c.Bind(&rurb)
	if err != nil {
		err := "incorrect body format"
		return rest_context.JSON(400, &RegisterUserResponse{
			Status: "error",
			Error:  &err,
			Data:   nil,
		})
	}

	if rurb.UserId == "" || rurb.Email == "" || rurb.Password == "" {
		err := "missing user id, email or password"
		return rest_context.JSON(400, &RegisterUserResponse{
			Status: "error",
			Error:  &err,
			Data:   nil,
		})
	}

	err = controllers.RegisterUser(&app_context, rurb.UserId, rurb.Email, rurb.Password)
	if err != nil {
		err := "invalid email or password"
		return rest_context.JSON(400, &RegisterUserResponse{
			Status: "error",
			Error:  &err,
			Data:   nil,
		})
	}

	return rest_context.JSON(200, &RegisterUserResponse{
		Status: "ok",
		Error:  nil,
		Data:   nil,
	})
}
