package transaction

import (
	echo "github.com/labstack/echo/v4"

	"go.a5r.dev/services/wallet/controllers"
	"go.a5r.dev/services/wallet/types"
)

type TransactionDepositRequestBody struct {
	Email string `json:"email"`
	Value int    `json:"value"`
}

type CreateTransactionDepositRouteResponse struct {
	Status string             `json:"status"`
	Error  *string            `json:"error"`
	Data   *types.Transaction `json:"data"`
}

func DepositRoute(c echo.Context) error {
	rest_context := c.(*types.RestContext)
	app_context := *(*rest_context).AppContext

	var tdrb TransactionDepositRequestBody
	err := rest_context.Bind(&tdrb)
	if err != nil {
		err := "missing or incorrect values for email or value body parameters"

		return rest_context.JSON(400, &CreateTransactionDepositRouteResponse{
			Status: "error",
			Error:  &err,
			Data:   nil,
		})
	}

	wallet, err := controllers.GetWalletController(&app_context, tdrb.Email)
	if err != nil {
		err := "failed to retrieve wallet"

		return rest_context.JSON(400, &CreateTransactionDepositRouteResponse{
			Status: "error",
			Error:  &err,
			Data:   nil,
		})
	}

	transaction_type := types.TransactionTypeDeposit
	transaction, err := controllers.CreateTransactionController(&app_context, tdrb.Email, transaction_type, tdrb.Value, wallet.CurrencyCode)
	if err != nil {
		err := err.Error()

		return rest_context.JSON(400, &CreateTransactionDepositRouteResponse{
			Status: "error",
			Error:  &err,
			Data:   nil,
		})
	}

	return rest_context.JSON(200, &CreateTransactionDepositRouteResponse{
		Status: "ok",
		Error:  nil,
		Data:   transaction,
	})
}
