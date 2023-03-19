package routes

import (
	"net/http"

	"github.com/autobar-dev/services/currency/controllers"
	"github.com/autobar-dev/services/currency/types"
	"github.com/labstack/echo"
)

type rateArgs struct {
	Base        string `query:"base" validate:"required"`
	Destination string `query:"destination" validate:"required"`
}

func Rate(c echo.Context) error {
	rc := c.(types.RestContext)

	var args rateArgs

	if err := rc.Bind(&args); err != nil {
		return rc.JSON(http.StatusBadRequest, &types.RestError{Error: err.Error()})
	}

	r, err := controllers.GetRate(args.Base, args.Destination, &rc.Stores.RateStore, &rc.Stores.SupportedCurrenciesStore, &rc.Stores.RemoteExchangeRateStore)

	if err != nil {
		return rc.JSON(http.StatusInternalServerError, &types.RestError{Error: err.Error()})
	}

	return rc.JSON(http.StatusOK, r)
}
