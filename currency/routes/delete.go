package routes

import (
	"net/http"

	"github.com/autobar-dev/services/currency/controllers"
	"github.com/autobar-dev/services/currency/types"
	"github.com/labstack/echo"
)

type restDeleteArgs struct {
	Code string `json:"code" form:"code"`
}

func Delete(c echo.Context) error {
	rc := c.(types.RestContext)

	var rda restDeleteArgs
	if err := rc.Bind(&rda); err != nil {
		return rc.JSON(http.StatusBadRequest, &types.RestError{Error: err.Error()})
	}

	dc, err := controllers.RemoveCurrency(rda.Code, &rc.Stores.SupportedCurrenciesStore)

	if err != nil {
		return rc.JSON(http.StatusBadRequest, &types.RestError{Error: err.Error()})
	}

	return rc.JSON(http.StatusOK, dc)
}
