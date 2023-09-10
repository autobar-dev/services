package transaction

import (
	"github.com/autobar-dev/services/wallet/controllers"
	"github.com/autobar-dev/services/wallet/types"
	"github.com/labstack/echo/v4"
)

type GetTransactionRouteResponse struct {
	Status string             `json:"status"`
	Error  *string            `json:"error"`
	Data   *types.Transaction `json:"data"`
}

func GetRoute(c echo.Context) error {
	rest_context := c.(*types.RestContext)
	app_context := *(*rest_context).AppContext

	id := c.QueryParam("id")

	if id == "" {
		err := "id query parameter not present"

		return rest_context.JSON(400, &GetTransactionRouteResponse{
			Status: "error",
			Error:  &err,
			Data:   nil,
		})
	}

	transaction, err := controllers.GetTransactionController(&app_context, id)
	if err != nil {
		err := err.Error()

		return rest_context.JSON(400, &GetTransactionRouteResponse{
			Status: "error",
			Error:  &err,
			Data:   nil,
		})
	}

	return rest_context.JSON(200, &GetTransactionRouteResponse{
		Status: "ok",
		Error:  nil,
		Data:   transaction,
	})
}
