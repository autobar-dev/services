package controllers

import (
	"errors"

	"go.a5r.dev/services/wallet/types"
)

func GetRateController(app_context *types.AppContext, from string, to string) (*types.Rate, error) {
	cr := app_context.Repositories.Currency

	to_currency, err := cr.GetCurrency(to)
	if err != nil {
		return nil, err
	}

	if !to_currency.Enabled {
		return nil, errors.New("'to' currency is disabled")
	}

	rate, err := cr.GetRate(from, to)
	if err != nil {
		return nil, err
	}

	return &types.Rate{
		From: rate.From,
		To:   rate.To,
		Rate: rate.Rate,
	}, nil
}
