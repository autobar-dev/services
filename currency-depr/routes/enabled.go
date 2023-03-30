package routes

import (
	"net/http"

	"github.com/autobar-dev/services/currency/controllers"
	"github.com/autobar-dev/services/currency/types"
	"github.com/labstack/echo"
)

// swagger:route GET /enabled Currency list-currencies
//
// Returns a list of enabled currencies and their details.
//
// responses:
//   200: enabledResponse
//   400: errorResponse

// List of all enabled currencies
//
// swagger:response enabledResponse
type _ struct {
	// in: body
	Body []types.SupportedCurrency
}

// Error response
//
// swagger:response errorResponse
type _ struct {
	// in: body
	Body types.RestError
}

func Enabled(c echo.Context) error {
	rc := c.(types.RestContext)

	sc, err := controllers.GetEnabledCurrencies(&rc.Stores.SupportedCurrenciesStore)

	if err != nil {
		return rc.JSON(http.StatusInternalServerError, &types.RestError{Error: "Could not fetch available currencies."})
	}

	return rc.JSON(http.StatusOK, sc)
}
