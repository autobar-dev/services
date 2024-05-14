package utils

import (
	"github.com/autobar-dev/services/promo/repositories"
	"github.com/autobar-dev/services/promo/types"
)

func PostgresPromoToPromo(pp *repositories.PostgresPromo) *types.Promo {
	var promo_type types.PromoType

	switch pp.Type {
	case "FIXED":
		promo_type = types.FixedType
	case "PERCENTAGE":
		promo_type = types.PercentageType
	}

	return &types.Promo{
		Id:        pp.Id,
		Type:      promo_type,
		ExpiresAt: pp.ExpiresAt,
		CreatedAt: pp.CreatedAt,
	}
}
