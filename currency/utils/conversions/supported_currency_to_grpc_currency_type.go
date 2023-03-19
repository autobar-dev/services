package conversions

import (
	"github.com/autobar-dev/services/currency/grpc/generated_grpc"
	"github.com/autobar-dev/services/currency/types"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func SupportedCurrencyToGrpcCurrencyType(sc *types.SupportedCurrency) *generated_grpc.CurrencyType {
	return &generated_grpc.CurrencyType{
		Code:      sc.Code,
		Name:      sc.Name,
		Enabled:   sc.Enabled,
		UpdatedAt: timestamppb.New(sc.UpdatedAt),
	}
}
