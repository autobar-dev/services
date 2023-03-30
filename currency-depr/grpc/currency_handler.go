package grpc

import (
	"github.com/autobar-dev/services/currency/grpc/generated_grpc"
	"github.com/autobar-dev/services/currency/types"
)

type CurrencyHandler struct {
	generated_grpc.UnimplementedCurrencyServer
	*types.AppContext
}

func NewCurrencyHandler(ac *types.AppContext) *CurrencyHandler {
	return &CurrencyHandler{
		AppContext:                  ac,
		UnimplementedCurrencyServer: generated_grpc.UnimplementedCurrencyServer{},
	}
}
