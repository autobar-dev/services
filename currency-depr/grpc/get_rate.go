package grpc

import (
	"context"

	"github.com/autobar-dev/services/currency/controllers"
	"github.com/autobar-dev/services/currency/grpc/generated_grpc"
	"github.com/autobar-dev/services/currency/utils/conversions"
)

func (ch *CurrencyHandler) GetRate(ctx context.Context, gr *generated_grpc.GetRateRequest) (*generated_grpc.RateType, error) {
	r, err := controllers.GetRate(gr.GetBase(), gr.GetDestination(), &ch.Stores.RateStore, &ch.Stores.SupportedCurrenciesStore, &ch.Stores.RemoteExchangeRateStore)

	if err != nil {
		return nil, err
	}

	return conversions.RateToGrpcRateType(r), nil
}
