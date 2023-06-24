package transaction

import (
	"math"

	echo "github.com/labstack/echo/v4"

	"go.a5r.dev/services/wallet/controllers"
	"go.a5r.dev/services/wallet/types"
)

type TransactionCurrencyChangeRequestBody struct {
	Email        string `json:"email"`
	CurrencyCode string `json:"currency_code"`
}

type CreateTransactionCurrencyChangeRouteResponse struct {
	Status string             `json:"status"`
	Error  *string            `json:"error"`
	Data   *types.Transaction `json:"data"`
}

func CurrencyChangeRoute(c echo.Context) error {
	rest_context := c.(*types.RestContext)
	app_context := *(*rest_context).AppContext

	var tccrb TransactionCurrencyChangeRequestBody
	err := rest_context.Bind(&tccrb)
	if err != nil {
		err := "missing or incorrect values for email or currency_code body parameters"

		return rest_context.JSON(400, &CreateTransactionCurrencyChangeRouteResponse{
			Status: "error",
			Error:  &err,
			Data:   nil,
		})
	}

	wallet, err := controllers.GetWalletController(&app_context, tccrb.Email)
	if err != nil {
		err := "failed to retrieve wallet"

		return rest_context.JSON(400, &CreateTransactionCurrencyChangeRouteResponse{
			Status: "error",
			Error:  &err,
			Data:   nil,
		})
	}

	rate, err := controllers.GetRateController(&app_context, wallet.CurrencyCode, tccrb.CurrencyCode)
	if err != nil {
		err := "failed to retrieve rate for specified currencies"

		return rest_context.JSON(400, &CreateTransactionCurrencyChangeRouteResponse{
			Status: "error",
			Error:  &err,
			Data:   nil,
		})
	}

	new_balance := math.Floor(float64(wallet.Balance) * rate.Rate)

	transaction_type := types.TransactionTypeCurrencyChange
	transaction, err := controllers.CreateTransactionController(&app_context, tccrb.Email, transaction_type, int(new_balance), tccrb.CurrencyCode)
	if err != nil {
		err := err.Error()

		return rest_context.JSON(400, &CreateTransactionCurrencyChangeRouteResponse{
			Status: "error",
			Error:  &err,
			Data:   nil,
		})
	}

	return rest_context.JSON(200, &CreateTransactionCurrencyChangeRouteResponse{
		Status: "ok",
		Error:  nil,
		Data:   transaction,
	})
}
