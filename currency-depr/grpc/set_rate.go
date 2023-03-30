package grpc

import (
	"context"

	"github.com/autobar-dev/services/currency/controllers"
	"github.com/autobar-dev/services/currency/grpc/generated_grpc"
	"github.com/autobar-dev/services/currency/utils/conversions"
)

func (ch *CurrencyHandler) SetRate(ctx context.Context, srr *generated_grpc.SetRateRequest) (*generated_grpc.RateType, error) {
	sr, err := controllers.SetRate(srr.GetBase(), srr.GetDestination(), srr.GetRate(), &ch.Stores.RateStore, &ch.Stores.SupportedCurrenciesStore)

	if err != nil {
		return nil, err
	}

	return conversions.RateToGrpcRateType(sr), nil
}
