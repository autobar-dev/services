package wallet

import (
	"github.com/labstack/echo/v4"
	"go.a5r.dev/services/wallet/controllers"
	"go.a5r.dev/services/wallet/types"
)

type GetWalletRouteResponse struct {
	Status string        `json:"status"`
	Error  *string       `json:"error"`
	Data   *types.Wallet `json:"data"`
}

func GetRoute(c echo.Context) error {
	rest_context := c.(*types.RestContext)
	app_context := *(*rest_context).AppContext

	user_email := c.QueryParam("email")

	if user_email == "" {
		err := "email query parameter not present"
		return rest_context.JSON(400, &GetWalletRouteResponse{
			Status: "error",
			Error:  &err,
			Data:   nil,
		})
	}

	wallet, err := controllers.GetWalletController(&app_context, user_email)
	if err != nil {
		err := err.Error()

		return rest_context.JSON(400, &GetWalletRouteResponse{
			Status: "error",
			Error:  &err,
			Data:   nil,
		})
	}

	return rest_context.JSON(400, &GetWalletRouteResponse{
		Status: "ok",
		Error:  nil,
		Data:   wallet,
	})
}
