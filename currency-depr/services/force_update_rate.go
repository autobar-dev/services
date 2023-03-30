package services

import "github.com/autobar-dev/services/currency/types/interfaces"

func ForceUpdateRate(base string, destination string, rs *interfaces.RateStore, rers *interfaces.RemoteExchangeRateStore) error {
	fer, err := FetchRemoteCurrencyRate(base, destination, rers)

	if err != nil {
		return err
	}

	return SetRate(fer.BaseCode, fer.DestinationCode, fer.ConversionRate, rs)
}
