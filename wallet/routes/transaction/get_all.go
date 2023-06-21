package transactions

import (
	"github.com/labstack/echo/v4"
	"go.a5r.dev/services/wallet/controllers"
	"go.a5r.dev/services/wallet/types"
)

type GetAllTransactionsRouteResponse struct {
	Status string               `json:"status"`
	Error  *string              `json:"error"`
	Data   *[]types.Transaction `json:"data"`
}

func GetAllRoute(c echo.Context) error {
	rest_context := c.(*types.RestContext)
	app_context := *(*rest_context).AppContext

	user_email := c.QueryParam("email")

	if user_email == "" {
		err := "email query parameter not present"
		return rest_context.JSON(400, &GetAllTransactionsRouteResponse{
			Status: "error",
			Error:  &err,
			Data:   nil,
		})
	}

	transactions, err := controllers.GetAllTransactionsController(&app_context, user_email)
	if err != nil {
		err := err.Error()

		return rest_context.JSON(400, &GetAllTransactionsRouteResponse{
			Status: "error",
			Error:  &err,
			Data:   nil,
		})
	}

	return rest_context.JSON(400, &GetAllTransactionsRouteResponse{
		Status: "ok",
		Error:  nil,
		Data:   transactions,
	})
}
