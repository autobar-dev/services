package routes

import (
	"net/http"

	"github.com/autobar-dev/services/currency/controllers"
	"github.com/autobar-dev/services/currency/types"
	"github.com/autobar-dev/services/currency/types/inputs"
	"github.com/labstack/echo"
)

type restCreateInput struct {
	Code string `json:"code" form:"code"`
	Name string `json:"name" form:"name"`
}

func Create(c echo.Context) error {
	rc := c.(types.RestContext)

	var i restCreateInput

	if err := rc.Bind(&i); err != nil {
		return rc.JSON(http.StatusBadRequest, &types.RestError{Error: "parameters not provided"})
	}

	inp := inputs.Currency(i)

	nc, err := controllers.CreateCurrency(&inp, &rc.Stores.SupportedCurrenciesStore)

	if err != nil {
		return rc.JSON(http.StatusBadRequest, err)
	}

	return rc.JSON(http.StatusCreated, nc)
}
