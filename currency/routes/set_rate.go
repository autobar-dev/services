package routes

import (
	"net/http"

	"github.com/autobar-dev/services/currency/controllers"
	"github.com/autobar-dev/services/currency/types"
	"github.com/labstack/echo"
)

type setRateArgs struct {
	Base        string  `json:"base" form:"base" validate:"required"`
	Destination string  `json:"destination" form:"destination" validate:"required"`
	Rate        float64 `json:"rate" form:"rate" validate:"required"`
}

func SetRate(c echo.Context) error {
	rc := c.(types.RestContext)

	var args setRateArgs

	if err := rc.Bind(&args); err != nil {
		return rc.JSON(http.StatusBadRequest, &types.RestError{Error: err.Error()})
	}

	sr, err := controllers.SetRate(args.Base, args.Destination, args.Rate, &rc.Stores.RateStore, &rc.Stores.SupportedCurrenciesStore)

	if err != nil {
		return rc.JSON(http.StatusInternalServerError, &types.RestError{Error: err.Error()})
	}

	return rc.JSON(http.StatusOK, sr)
}
