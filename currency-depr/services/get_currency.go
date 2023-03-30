package services

import "github.com/autobar-dev/services/currency/types/interfaces"

func GetCurrency(currency_code string, scs *interfaces.SupportedCurrenciesStore) (*interfaces.SupportedCurrenciesStoreRow, error) {
	return (*scs).Get(currency_code)
}
