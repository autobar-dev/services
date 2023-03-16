package controllers

import (
	"errors"

	"github.com/autobar-dev/services/currency/services"
	"github.com/autobar-dev/services/currency/types"
	"github.com/autobar-dev/services/currency/types/interfaces"
	"github.com/autobar-dev/services/currency/utils/conversions"
)

func GetSupportedCurrencies(scs *interfaces.SupportedCurrenciesStore) (*[]types.SupportedCurrency, error) {
	scr, err := services.GetSupportedCurrencies(scs)

	if err != nil {
		return nil, errors.New("could not fetch supported currencies")
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
