package wallet

import (
	"github.com/labstack/echo/v4"
	"go.a5r.dev/services/wallet/types"
)

type GetWalletRouteResponse struct {
	Status string        `json:"status"`
	Error  *string       `json:"error"`
	Data   *types.Wallet `json:"data"`
}

func GetRoute(c echo.Context) error {
	rest_context := c.(*types.RestContext)
	_ = *(*rest_context).AppContext

	return nil
}
