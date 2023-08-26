package routes

import (
	"github.com/autobar-dev/services/auth/controllers"
	"github.com/autobar-dev/services/auth/types"
	"github.com/labstack/echo/v4"
)

type LoginModuleRequestBody struct {
	SerialNumber string `json:"serial_number"`
	PrivateKey   string `json:"private_key"`
}

type LoginModuleResponse struct {
	Status string        `json:"status"`
	Error  *string       `json:"error"`
	Data   *types.Tokens `json:"data"`
}

func LoginModuleRoute(c echo.Context) error {
	rest_context := c.(*types.RestContext)
	app_context := *(*rest_context).AppContext

	var lmrb LoginModuleRequestBody
	err := c.Bind(&lmrb)
	if err != nil {
		err := "incorrect body format"
		return rest_context.JSON(400, &LoginModuleResponse{
			Status: "error",
			Error:  &err,
			Data:   nil,
		})
	}

	if lmrb.SerialNumber == "" || lmrb.PrivateKey == "" {
		err := "missing serial number or private key"
		return rest_context.JSON(400, &LoginModuleResponse{
			Status: "error",
			Error:  &err,
			Data:   nil,
		})
	}

	tokens, err := controllers.LoginModule(&app_context, lmrb.SerialNumber, lmrb.PrivateKey)
	if err != nil {
		err := "incorrect serial number or private key"
		return rest_context.JSON(400, &LoginModuleResponse{
			Status: "error",
			Error:  &err,
			Data:   nil,
		})
	}

	return rest_context.JSON(200, &LoginModuleResponse{
		Status: "ok",
		Error:  nil,
		Data:   tokens,
	})
}
