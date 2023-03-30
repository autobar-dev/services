package conversions

import (
	"github.com/autobar-dev/services/currency/types"
	"github.com/autobar-dev/services/currency/types/interfaces"
)

func SupportedCurrenciesStoreRowToSupportedCurrency(scsr *interfaces.SupportedCurrenciesStoreRow) *types.SupportedCurrency {
	return &types.SupportedCurrency{
		Code:      scsr.Code,
		Name:      scsr.Name,
		Enabled:   scsr.Enabled,
		UpdatedAt: scsr.UpdatedAt,
	}
}
