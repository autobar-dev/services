package utils

import (
	"go.a5r.dev/services/wallet/repositories"
	"go.a5r.dev/services/wallet/types"
)

func ServiceRateToRate(sr repositories.ServiceRate) *types.Rate {
	return &types.Rate{
		From: sr.From,
		To:   sr.To,
		Rate: sr.Rate,
	}
}
