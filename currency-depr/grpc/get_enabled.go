package grpc

import (
	"context"

	"github.com/autobar-dev/services/currency/controllers"
	"github.com/autobar-dev/services/currency/grpc/generated_grpc"
	"github.com/autobar-dev/services/currency/utils/conversions"
)

func (ch *CurrencyHandler) GetSupported(ctx context.Context, _ *generated_grpc.GetSupportedRequest) (*generated_grpc.GetSupportedResponse, error) {
	cl, err := controllers.GetEnabledCurrencies(&ch.Stores.SupportedCurrenciesStore)

	if err != nil {
		return nil, err
	}

	gsr := generated_grpc.GetSupportedResponse{
		Currencies: []*generated_grpc.CurrencyType{},
	}

	for _, currency := range *cl {
		gsr.Currencies = append(gsr.Currencies, conversions.SupportedCurrencyToGrpcCurrencyType(&currency))
	}

	return &gsr, nil
}
