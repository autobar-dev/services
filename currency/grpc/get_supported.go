package grpc

import (
	"context"

	"github.com/autobar-dev/services/currency/controllers"
	"github.com/autobar-dev/services/currency/grpc/generated_grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (ch *CurrencyHandler) GetSupported(ctx context.Context, ar *generated_grpc.GetSupportedRequest) (*generated_grpc.GetSupportedResponse, error) {
	cl, err := controllers.GetSupportedCurrencies(&ch.Stores.SupportedCurrenciesStore)

	if err != nil {
		return nil, err
	}

	gsr := generated_grpc.GetSupportedResponse{
		Currencies: []*generated_grpc.CurrencyType{},
	}

	for _, currency := range *cl {
		gsr.Currencies = append(gsr.Currencies, &generated_grpc.CurrencyType{
			Code:      currency.Code,
			Name:      currency.Name,
			Enabled:   currency.Enabled,
			UpdatedAt: timestamppb.New(currency.UpdatedAt),
		})
	}

	return &gsr, nil
}
