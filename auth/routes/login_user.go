package routes

import (
	"github.com/autobar-dev/services/auth/controllers"
	"github.com/autobar-dev/services/auth/types"
	"github.com/labstack/echo/v4"
)

type LoginUserRequestBody struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	RememberMe bool   `json:"remember_me"`
}

type LoginUserResponse struct {
	Status string        `json:"status"`
	Error  *string       `json:"error"`
	Data   *types.Tokens `json:"data"`
}

func LoginUserRoute(c echo.Context) error {
	rest_context := c.(*types.RestContext)
	app_context := *(*rest_context).AppContext

	var lurb LoginUserRequestBody
	err := c.Bind(&lurb)
	if err != nil {
		err := "incorrect body format"
		return rest_context.JSON(400, &LoginUserResponse{
			Status: "error",
			Error:  &err,
			Data:   nil,
		})
	}

	if lurb.Email == "" || lurb.Password == "" {
		err := "missing email or password"
		return rest_context.JSON(400, &LoginUserResponse{
			Status: "error",
			Error:  &err,
			Data:   nil,
		})
	}

	tokens, err := controllers.LoginUser(&app_context, lurb.Email, lurb.Password, lurb.RememberMe)
	if err != nil {
		err := "incorrect email or password"
		return rest_context.JSON(400, &LoginUserResponse{
			Status: "error",
			Error:  &err,
			Data:   nil,
		})
	}

	return rest_context.JSON(200, &LoginUserResponse{
		Status: "ok",
		Error:  nil,
		Data:   tokens,
	})
}
