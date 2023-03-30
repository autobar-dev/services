package grpc

import (
	"context"

	"github.com/autobar-dev/services/currency/controllers"
	"github.com/autobar-dev/services/currency/grpc/generated_grpc"
	"github.com/autobar-dev/services/currency/utils/conversions"
)

func (ch *CurrencyHandler) SetEnabled(ctx context.Context, ser *generated_grpc.SetEnabledRequest) (*generated_grpc.CurrencyType, error) {
	ec, err := controllers.SetEnabledCurrency(ser.GetCode(), ser.GetEnabled(), &ch.Stores.SupportedCurrenciesStore)

	if err != nil {
		return nil, err
	}

	return conversions.SupportedCurrencyToGrpcCurrencyType(ec), nil
}
