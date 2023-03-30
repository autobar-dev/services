package grpc

import (
	"context"

	"github.com/autobar-dev/services/currency/controllers"
	"github.com/autobar-dev/services/currency/grpc/generated_grpc"
	"github.com/autobar-dev/services/currency/utils/conversions"
)

func (ch *CurrencyHandler) Delete(ctx context.Context, dr *generated_grpc.DeleteRequest) (*generated_grpc.CurrencyType, error) {
	dc, err := controllers.RemoveCurrency(dr.GetCode(), &ch.Stores.SupportedCurrenciesStore)

	if err != nil {
		return nil, err
	}

	return conversions.SupportedCurrencyToGrpcCurrencyType(dc), nil
}
