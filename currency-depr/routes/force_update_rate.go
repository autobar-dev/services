package routes

import (
	"net/http"

	"github.com/autobar-dev/services/currency/controllers"
	"github.com/autobar-dev/services/currency/types"
	"github.com/labstack/echo"
)

type forceUpdateRateArgs struct {
	Base        string `json:"base" form:"base" validate:"required"`
	Destination string `json:"destination" form:"destination" validate:"required"`
}

func ForceUpdateRate(c echo.Context) error {
	rc := c.(types.RestContext)

	var args forceUpdateRateArgs

	if err := rc.Bind(&args); err != nil {
		return rc.JSON(http.StatusBadRequest, &types.RestError{Error: err.Error()})
	}

	ur, err := controllers.ForceUpdateRate(args.Base, args.Destination, &rc.Stores.RateStore, &rc.Stores.SupportedCurrenciesStore, &rc.Stores.RemoteExchangeRateStore)

	if err != nil {
		return rc.JSON(http.StatusInternalServerError, &types.RestError{Error: err.Error()})
	}

	return rc.JSON(http.StatusOK, ur)
}
