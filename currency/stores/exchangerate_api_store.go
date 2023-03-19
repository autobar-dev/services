package stores

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/autobar-dev/services/currency/types/interfaces"
)

type ExchangeRateApiStore struct{}

type exchangeRateApiResponse struct {
	Result             string  `json:"result"`
	Documentation      string  `json:"documentation"`
	TermsOfUse         string  `json:"terms_of_use"`
	TimeLastUpdateUnix int     `json:"time_last_update_unix"`
	TimeLastUpdateUtc  string  `json:"time_last_update_utc"`
	TimeNextUpdateUnix int     `json:"time_next_update_unix"`
	TimeNextUpdateUtc  string  `json:"time_next_update_utc"`
	BaseCode           string  `json:"base_code"`
	TargetCode         string  `json:"target_code"`
	ConversionRate     float64 `json:"conversion_rate"`
}

func NewExchangeRateApiStore() (*ExchangeRateApiStore, error) {
	return &ExchangeRateApiStore{}, nil
}

func (era *ExchangeRateApiStore) Get(base string, destination string) (*interfaces.RemoteExchangeRateStoreRow, error) {
	erak := os.Getenv("EXCHANGE_RATE_API_KEY")

	resp, err := http.Get(
		fmt.Sprintf(
			"https://v6.exchangerate-api.com/v6/%s/pair/%s/%s",
			erak,
			base,
			destination,
		),
	)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var fer exchangeRateApiResponse
	err = json.NewDecoder(resp.Body).Decode(&fer)

	if err != nil {
		return nil, err
	}

	return &interfaces.RemoteExchangeRateStoreRow{
		BaseCode:        fer.BaseCode,
		DestinationCode: fer.TargetCode,
		ConversionRate:  fer.ConversionRate,
	}, nil
}
