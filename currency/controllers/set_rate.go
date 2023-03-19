package controllers

import (
	"errors"
	"strings"

	"github.com/autobar-dev/services/currency/services"
	"github.com/autobar-dev/services/currency/types"
	"github.com/autobar-dev/services/currency/types/interfaces"
	"github.com/autobar-dev/services/currency/utils/conversions"
)

func SetRate(base string, destination string, rate float64, rs *interfaces.RateStore, scs *interfaces.SupportedCurrenciesStore) (*types.Rate, error) {
	base = strings.ToUpper(base)
	destination = strings.ToUpper(destination)

	base_currency, b_err := services.GetCurrency(base, scs)
	dest_currency, d_err := services.GetCurrency(destination, scs)

	if b_err != nil || d_err != nil {
		return nil, errors.New("one or both of the provided currencies is not supported")
	}

	if !base_currency.Enabled || !dest_currency.Enabled {
		return nil, errors.New("one or both of the provided currencies is not enabled")
	}

	if rate == 0 {
		return nil, errors.New("rate cannot be zero")
	}

	err := services.SetRate(base, destination, rate, rs)

	if err != nil {
		return nil, err
	}

	r, err := services.GetRate(base, destination, rs)

	if err != nil {
		return nil, err
	}

	return conversions.RateStoreRowToRate(r), nil
}
