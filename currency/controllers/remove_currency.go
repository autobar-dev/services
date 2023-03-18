package controllers

import (
	"strings"

	"github.com/autobar-dev/services/currency/services"
	"github.com/autobar-dev/services/currency/types"
	"github.com/autobar-dev/services/currency/types/interfaces"
	"github.com/autobar-dev/services/currency/utils/conversions"
)

func RemoveCurrency(currency_code string, scs *interfaces.SupportedCurrenciesStore) (*types.SupportedCurrency, error) {
	currency_code = strings.ToUpper(currency_code)

	c, err := services.GetCurrency(currency_code, scs)

	if err != nil {
		return nil, err
	}

	err = services.RemoveCurrency(currency_code, scs)

	if err != nil {
		return nil, err
	}

	return conversions.SupportedCurrenciesStoreRowToSupportedCurrency(c), nil
}
