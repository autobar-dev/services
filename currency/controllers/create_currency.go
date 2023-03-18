package controllers

import (
	"errors"
	"strings"

	"github.com/autobar-dev/services/currency/services"
	"github.com/autobar-dev/services/currency/types"
	"github.com/autobar-dev/services/currency/types/inputs"
	"github.com/autobar-dev/services/currency/types/interfaces"
	"github.com/autobar-dev/services/currency/utils/conversions"
)

func CreateCurrency(input *inputs.Currency, scs *interfaces.SupportedCurrenciesStore) (*types.SupportedCurrency, error) {
	// Check validity
	if len(input.Code) != 3 {
		return nil, errors.New("currency code has to be 3 letters long")
	}

	if len(input.Name) == 0 {
		return nil, errors.New("currency needs to have a name")
	}

	// Check if already exists
	c, _ := services.GetCurrency(input.Code, scs)

	if c != nil {
		return nil, errors.New("currency with this code already exists")
	}

	input.Code = strings.ToUpper(input.Code)

	err := services.CreateCurrency(input, scs)

	if err != nil {
		return nil, err
	}

	nc, err := services.GetCurrency(input.Code, scs)

	if err != nil {
		return nil, err
	}

	return conversions.SupportedCurrenciesStoreRowToSupportedCurrency(nc), nil
}
