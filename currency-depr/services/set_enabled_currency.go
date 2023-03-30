package services

import (
	"github.com/autobar-dev/services/currency/types/interfaces"
)

func SetEnabledCurrency(currency_code string, enabled bool, supported_currencies_store *interfaces.SupportedCurrenciesStore) error {
	return (*supported_currencies_store).SetEnabled(currency_code, enabled)
}
