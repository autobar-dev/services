package routes

import (
	"github.com/autobar-dev/services/auth/controllers"
	"github.com/autobar-dev/services/auth/types"
	"github.com/labstack/echo/v4"
)

type RegisterModuleRequestBody struct {
	SerialNumber string `json:"serial_number"`
}

type RegisterModuleResponseData struct {
	PrivateKey string `json:"private_key"`
}

type RegisterModuleResponse struct {
	Status string                      `json:"status"`
	Error  *string                     `json:"error"`
	Data   *RegisterModuleResponseData `json:"data"`
}

func RegisterModuleRoute(c echo.Context) error {
	rest_context := c.(*types.RestContext)
	app_context := *(*rest_context).AppContext

	var rmrb RegisterModuleRequestBody
	err := c.Bind(&rmrb)
	if err != nil {
		err := "incorrect body format"
		return rest_context.JSON(400, &RegisterModuleResponse{
			Status: "error",
			Error:  &err,
			Data:   nil,
		})
	}

	if rmrb.SerialNumber == "" {
		err := "missing serial number"
		return rest_context.JSON(400, &RegisterModuleResponse{
			Status: "error",
			Error:  &err,
			Data:   nil,
		})
	}

	private_key, err := controllers.RegisterModule(&app_context, rmrb.SerialNumber)
	if err != nil {
		err := "invalid serial number"
		return rest_context.JSON(400, &RegisterModuleResponse{
			Status: "error",
			Error:  &err,
			Data:   nil,
		})
	}

	return rest_context.JSON(200, &RegisterModuleResponse{
		Status: "ok",
		Error:  nil,
		Data: &RegisterModuleResponseData{
			PrivateKey: *private_key,
		},
	})
}
