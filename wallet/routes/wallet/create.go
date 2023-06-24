package wallet

import (
	echo "github.com/labstack/echo/v4"
	"go.a5r.dev/services/wallet/controllers"
	"go.a5r.dev/services/wallet/types"
)

type CreateWalletRequestBody struct {
	Email        string `json:"email"`
	CurrencyCode string `json:"currency_code"`
}

type CreateWalletRouteResponse struct {
	Status string        `json:"status"`
	Error  *string       `json:"error"`
	Data   *types.Wallet `json:"data"`
}

func CreateRoute(c echo.Context) error {
	rest_context := c.(*types.RestContext)
	app_context := *(*rest_context).AppContext

	var cwrb CreateWalletRequestBody
	err := rest_context.Bind(&cwrb)
	if err != nil {
		err := "missing or incorrect values for email or currency_code body parameters"

		return rest_context.JSON(400, &CreateWalletRouteResponse{
			Status: "error",
			Error:  &err,
			Data:   nil,
		})
	}

	wallet, err := controllers.CreateWalletController(&app_context, cwrb.Email, cwrb.CurrencyCode)
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
		Data:   wallet,
	})
}
