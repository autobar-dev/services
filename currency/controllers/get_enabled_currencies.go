package controllers

import (
	"errors"

	"github.com/autobar-dev/services/currency/services"
	"github.com/autobar-dev/services/currency/types"
	"github.com/autobar-dev/services/currency/types/interfaces"
	"github.com/autobar-dev/services/currency/utils/conversions"
)

func GetEnabledCurrencies(scs *interfaces.SupportedCurrenciesStore) (*[]types.SupportedCurrency, error) {
	scr, err := services.GetEnabledCurrencies(scs)

	if err != nil {
		return nil, errors.New("could not fetch enabled currencies")
	}

	sc := []types.SupportedCurrency{}

	for _, s := range *scr {
		sc = append(sc,
			*conversions.SupportedCurrenciesStoreRowToSupportedCurrency(
				&s,
			),
		)
	}

	return &sc, err
}
