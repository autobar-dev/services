package wallet

import (
	echo "github.com/labstack/echo/v4"
	"go.a5r.dev/services/wallet/controllers"
	"go.a5r.dev/services/wallet/types"
)

type CreateWalletRouteResponse struct {
	Status string      `json:"status"`
	Error  *string     `json:"error"`
	Data   interface{} `json:"data"`
}

func CreateRoute(c echo.Context) error {
	rest_context := c.(*types.RestContext)
	app_context := *(*rest_context).AppContext

	err := controllers.CreateWalletController(&app_context)
	if err != nil {
		err := err.Error()

		return rest_context.JSON(400, &CreateWalletRouteResponse{
			Status: "error",
			Error:  &err,
			Data:   nil,
		})
	}

	return rest_context.JSON(200, &CreateWalletRouteResponse{
		Status: "ok",
		Error:  nil,
		Data:   nil,
	})
}
