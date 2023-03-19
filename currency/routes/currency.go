package routes

import (
	"net/http"

	"github.com/autobar-dev/services/currency/controllers"
	"github.com/autobar-dev/services/currency/types"
	"github.com/labstack/echo"
)

type restGetCurrencyArgs struct {
	Code string `query:"code" validate:"required"`
}

func Currency(c echo.Context) error {
	rc := c.(types.RestContext)

	var rgca restGetCurrencyArgs

	if err := rc.Bind(&rgca); err != nil {
		return rc.JSON(http.StatusBadRequest, err.Error())
	}

	gc, err := controllers.GetCurrency(rgca.Code, &rc.Stores.SupportedCurrenciesStore)

	if err != nil {
		return rc.JSON(http.StatusBadRequest, &types.RestError{Error: err.Error()})
	}

	return rc.JSON(http.StatusOK, gc)
}
