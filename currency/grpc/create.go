package grpc

import (
	"context"

	"github.com/autobar-dev/services/currency/controllers"
	"github.com/autobar-dev/services/currency/grpc/generated_grpc"
	"github.com/autobar-dev/services/currency/types/inputs"
	"github.com/autobar-dev/services/currency/utils/conversions"
)

func (ch *CurrencyHandler) Create(ctx context.Context, cr *generated_grpc.CreateRequest) (*generated_grpc.CurrencyType, error) {
	cc, err := controllers.CreateCurrency(&inputs.Currency{
		Code: cr.Code,
		Name: cr.Name,
	}, &ch.Stores.SupportedCurrenciesStore)

	if err != nil {
		return nil, err
	}

	return conversions.SupportedCurrencyToGrpcCurrencyType(cc), nil
}
