package grpc

import (
	"context"

	"github.com/autobar-dev/services/currency/controllers"
	"github.com/autobar-dev/services/currency/grpc/generated_grpc"
	"github.com/autobar-dev/services/currency/utils/conversions"
)

func (ch *CurrencyHandler) ForceUpdateRate(ctx context.Context, fur *generated_grpc.ForceUpdateRateRequest) (*generated_grpc.RateType, error) {
	r, err := controllers.ForceUpdateRate(fur.GetBase(), fur.GetDestination(), &ch.Stores.RateStore, &ch.Stores.SupportedCurrenciesStore, &ch.Stores.RemoteExchangeRateStore)

	if err != nil {
		(*ch.AppLogger).Error("ForceUpdateRate error", err)
		return nil, err
	}

	return conversions.RateToGrpcRateType(r), nil
}
