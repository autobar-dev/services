package wallet

import (
	"github.com/autobar-dev/services/wallet/controllers"
	"github.com/autobar-dev/services/wallet/types"
	echo "github.com/labstack/echo/v4"
)

type CreateWalletRequestBody struct {
	UserId       string `json:"user_id"`
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
		err := "missing or incorrect values for user_id or currency_code body parameters"

		return rest_context.JSON(400, &CreateWalletRouteResponse{
			Status: "error",
			Error:  &err,
			Data:   nil,
		})
	}

	wallet, err := controllers.CreateWalletController(&app_context, cwrb.UserId, cwrb.CurrencyCode)
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
