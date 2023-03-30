package grpc

import (
	"context"

	"github.com/autobar-dev/services/currency/controllers"
	"github.com/autobar-dev/services/currency/grpc/generated_grpc"
	"github.com/autobar-dev/services/currency/utils/conversions"
)

func (ch *CurrencyHandler) Get(ctx context.Context, gr *generated_grpc.GetRequest) (*generated_grpc.CurrencyType, error) {
	gc, err := controllers.GetCurrency(gr.GetCode(), &ch.Stores.SupportedCurrenciesStore)

	if err != nil {
		return nil, err
	}

	return conversions.SupportedCurrencyToGrpcCurrencyType(gc), nil
}
