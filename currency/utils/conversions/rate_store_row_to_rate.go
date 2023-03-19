package conversions

import (
	"github.com/autobar-dev/services/currency/types"
	"github.com/autobar-dev/services/currency/types/interfaces"
)

func RateStoreRowToRate(rsr *interfaces.RateStoreRow) *types.Rate {
	return &types.Rate{
		BaseCurrency:        rsr.BaseCurrency,
		DestinationCurrency: rsr.DestinationCurrency,
		Rate:                rsr.Rate,
		UpdatedAt:           rsr.UpdatedAt,
	}
}
