package services

import (
	"github.com/autobar-dev/services/currency/types/interfaces"
)

func FetchRemoteCurrencyRate(base string, destination string, rers *interfaces.RemoteExchangeRateStore) (*interfaces.RemoteExchangeRateStoreRow, error) {
	fer, err := (*rers).Get(base, destination)

	if err != nil {
		return nil, err
	}

	return fer, nil
}
