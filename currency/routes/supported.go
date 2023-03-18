package routes

import (
	"net/http"

	"github.com/autobar-dev/services/currency/controllers"
	"github.com/autobar-dev/services/currency/types"
	"github.com/labstack/echo"
)

func Supported(c echo.Context) error {
	rc := c.(types.RestContext)

	sc, err := controllers.GetSupportedCurrencies(&rc.Stores.SupportedCurrenciesStore)

	if err != nil {
		return rc.JSON(http.StatusInternalServerError, &types.RestError{Error: "Could not fetch available currencies."})
	}

	return rc.JSON(http.StatusOK, sc)
}
