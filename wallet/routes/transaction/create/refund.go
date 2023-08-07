package transaction

import (
	echo "github.com/labstack/echo/v4"

	"go.a5r.dev/services/wallet/controllers"
	"go.a5r.dev/services/wallet/types"
)

type TransactionRefundRequestBody struct {
	UserId string `json:"user_id"`
	Value  int    `json:"value"`
}

type CreateTransactionRefundRouteResponse struct {
	Status string             `json:"status"`
	Error  *string            `json:"error"`
	Data   *types.Transaction `json:"data"`
}

func RefundRoute(c echo.Context) error {
	rest_context := c.(*types.RestContext)
	app_context := *(*rest_context).AppContext

	var trrb TransactionRefundRequestBody
	err := rest_context.Bind(&trrb)
	if err != nil {
		err := "missing or incorrect values for user_id or value body parameters"

		return rest_context.JSON(400, &CreateTransactionRefundRouteResponse{
			Status: "error",
			Error:  &err,
			Data:   nil,
		})
	}

	wallet, err := controllers.GetWalletController(&app_context, trrb.UserId)
	if err != nil {
		err := "failed to retrieve wallet"

		return rest_context.JSON(400, &CreateTransactionRefundRouteResponse{
			Status: "error",
			Error:  &err,
			Data:   nil,
		})
	}

	transaction_type := types.TransactionTypeRefund
	transaction, err := controllers.CreateTransactionController(&app_context, trrb.UserId, transaction_type, trrb.Value, wallet.CurrencyCode)
	if err != nil {
		err := err.Error()

		return rest_context.JSON(400, &CreateTransactionRefundRouteResponse{
			Status: "error",
			Error:  &err,
			Data:   nil,
		})
	}

	return rest_context.JSON(200, &CreateTransactionRefundRouteResponse{
		Status: "ok",
		Error:  nil,
		Data:   transaction,
	})
}
