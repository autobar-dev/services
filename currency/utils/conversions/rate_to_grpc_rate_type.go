package conversions

import (
	"github.com/autobar-dev/services/currency/grpc/generated_grpc"
	"github.com/autobar-dev/services/currency/types"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func RateToGrpcRateType(r *types.Rate) *generated_grpc.RateType {
	return &generated_grpc.RateType{
		Base:        r.BaseCurrency,
		Destination: r.DestinationCurrency,
		Rate:        r.Rate,
		UpdatedAt:   timestamppb.New(r.UpdatedAt),
	}
}
