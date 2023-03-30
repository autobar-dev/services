package services

import "github.com/autobar-dev/services/currency/types/interfaces"

func RemoveCurrency(currency_code string, scs *interfaces.SupportedCurrenciesStore) error {
	return (*scs).Delete(currency_code)
}
