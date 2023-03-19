package services

import "github.com/autobar-dev/services/currency/types/interfaces"

func GetEnabledCurrencies(scs *interfaces.SupportedCurrenciesStore) (*[]interfaces.SupportedCurrenciesStoreRow, error) {
	return (*scs).GetAll()
}
