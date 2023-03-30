package controllers

import (
	"errors"
	"strings"

	"github.com/autobar-dev/services/currency/services"
	"github.com/autobar-dev/services/currency/types"
	"github.com/autobar-dev/services/currency/types/interfaces"
	"github.com/autobar-dev/services/currency/utils/conversions"
)

func SetEnabledCurrency(currency_code string, enabled bool, scs *interfaces.SupportedCurrenciesStore) (*types.SupportedCurrency, error) {
	currency_code = strings.ToUpper(currency_code)

	c, _ := services.GetCurrency(currency_code, scs)

	if c == nil {
		return nil, errors.New("currency to enable/disable does not exist")
	}

	err := services.SetEnabledCurrency(currency_code, enabled, scs)

	if err != nil {
		return nil, err
	}

	nc, err := services.GetCurrency(currency_code, scs)

	if err != nil {
		return nil, err
	}

	return conversions.SupportedCurrenciesStoreRowToSupportedCurrency(nc), nil
}
