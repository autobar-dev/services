package transaction

import (
	"github.com/autobar-dev/services/wallet/controllers"
	"github.com/autobar-dev/services/wallet/types"
	"github.com/labstack/echo/v4"
)

type GetAllTransactionsRouteResponse struct {
	Status string               `json:"status"`
	Error  *string              `json:"error"`
	Data   *[]types.Transaction `json:"data"`
}

func GetAllRoute(c echo.Context) error {
	rest_context := c.(*types.RestContext)
	app_context := *(*rest_context).AppContext
	client_context := rest_context.ClientContext

	var user_id string

	user_id = c.QueryParam("user_id")
	if user_id == "" {
		if client_context != nil {
			user_id = client_context.Identifier
		} else {
			err := "user_id query parameter not present nor authenticated"
			return rest_context.JSON(400, &GetAllTransactionsRouteResponse{
				Status: "error",
				Error:  &err,
				Data:   nil,
			})
		}
	}

	transactions, err := controllers.GetAllTransactionsController(&app_context, user_id)
	if err != nil {
		err := err.Error()

		return rest_context.JSON(404, &GetAllTransactionsRouteResponse{
			Status: "error",
			Error:  &err,
			Data:   nil,
		})
	}

	return rest_context.JSON(200, &GetAllTransactionsRouteResponse{
		Status: "ok",
		Error:  nil,
		Data:   transactions,
	})
}
