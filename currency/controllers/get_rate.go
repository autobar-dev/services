package controllers

import (
	"errors"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/autobar-dev/services/currency/services"
	"github.com/autobar-dev/services/currency/types"
	"github.com/autobar-dev/services/currency/types/interfaces"
	"github.com/autobar-dev/services/currency/utils/conversions"
)

func GetRate(base_code string, destination_code string, rs *interfaces.RateStore, scs *interfaces.SupportedCurrenciesStore, rers *interfaces.RemoteExchangeRateStore) (*types.Rate, error) {
	base_code = strings.ToUpper(base_code)
	destination_code = strings.ToUpper(destination_code)

	base_currency, b_err := services.GetCurrency(base_code, scs)
	dest_currency, d_err := services.GetCurrency(destination_code, scs)

	if b_err != nil || d_err != nil {
		return nil, errors.New("one or both of the provided currencies is not supported")
	}

	if !base_currency.Enabled || !dest_currency.Enabled {
		return nil, errors.New("one or both of the provided currencies is not enabled")
	}

	r, err := services.GetRate(base_code, destination_code, rs)

	if err != nil {
		if err.Error() != "sql: no rows in result set" {
			return nil, err
		}

		err = services.ForceUpdateRate(base_code, destination_code, rs, rers)

		if err != nil {
			return nil, err
		}

		r, err = services.GetRate(base_code, destination_code, rs)

		if err != nil {
			return nil, err
		}
	}

	now := time.Now()
	prr_int, _ := strconv.Atoi(os.Getenv("PAST_RATE_RETENTION"))
	past_rate_retention := float64(prr_int)

	if now.Sub(r.UpdatedAt).Seconds() > past_rate_retention {
		err = services.ForceUpdateRate(base_code, destination_code, rs, rers)

		if err != nil {
			return nil, err
		}

		r, err = services.GetRate(base_code, destination_code, rs)

		if err != nil {
			return nil, err
		}
	}

	return conversions.RateStoreRowToRate(r), nil
}
