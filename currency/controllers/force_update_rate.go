package controllers

import (
	"errors"
	"strings"

	"github.com/autobar-dev/services/currency/services"
	"github.com/autobar-dev/services/currency/types"
	"github.com/autobar-dev/services/currency/types/interfaces"
	"github.com/autobar-dev/services/currency/utils/conversions"
)

func ForceUpdateRate(base string, dest string, rs *interfaces.RateStore, scs *interfaces.SupportedCurrenciesStore, rers *interfaces.RemoteExchangeRateStore) (*types.Rate, error) {
	base = strings.ToUpper(base)
	dest = strings.ToUpper(dest)

	base_currency, b_err := services.GetCurrency(base, scs)
	dest_currency, d_err := services.GetCurrency(dest, scs)

	if b_err != nil || d_err != nil {
		return nil, errors.New("one or both of the provided currencies is not supported")
	}

	if !base_currency.Enabled || !dest_currency.Enabled {
		return nil, errors.New("one or both of the provided currencies is not enabled")
	}

	err := services.ForceUpdateRate(base, dest, rs, rers)

	if err != nil {
		return nil, err
	}

	rate, err := services.GetRate(base, dest, rs)

	if err != nil {
		return nil, err
	}

	return conversions.RateStoreRowToRate(rate), nil
}
