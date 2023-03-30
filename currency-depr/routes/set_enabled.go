package routes

import (
	"net/http"

	"github.com/autobar-dev/services/currency/controllers"
	"github.com/autobar-dev/services/currency/types"
	"github.com/labstack/echo"
)

type restSetEnabledInput struct {
	Code    string `json:"code" form:"code" validate:"required"`
	Enabled bool   `json:"enabled" form:"enabled" validate:"required"`
}

func SetEnabled(c echo.Context) error {
	rc := c.(types.RestContext)

	var i restSetEnabledInput

	if err := rc.Bind(&i); err != nil {
		return rc.JSON(http.StatusBadRequest, &types.RestError{Error: "parameters either incorrect or not provided"})
	}

	ec, err := controllers.SetEnabledCurrency(i.Code, i.Enabled, &rc.Stores.SupportedCurrenciesStore)

	if err != nil {
		return rc.JSON(http.StatusBadRequest, &types.RestError{Error: err.Error()})
	}

	return c.JSON(http.StatusOK, ec)
}
