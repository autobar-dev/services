package services

import (
	"github.com/autobar-dev/services/currency/types/inputs"
	"github.com/autobar-dev/services/currency/types/interfaces"
)

func CreateCurrency(input *inputs.Currency, scs *interfaces.SupportedCurrenciesStore) error {
	return (*scs).Insert(input)
}
