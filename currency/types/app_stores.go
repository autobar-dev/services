package types

import "github.com/autobar-dev/services/currency/types/interfaces"

type AppStores struct {
	RateStore                interfaces.RateStore
	SupportedCurrenciesStore interfaces.SupportedCurrenciesStore
}
