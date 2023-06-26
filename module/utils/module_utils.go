package utils

import (
	"encoding/json"

	"go.a5r.dev/services/module/repositories"
	"go.a5r.dev/services/module/types"
)

func PostgresModuleToModule(pm repositories.PostgresModule) *types.Module {
	var prices map[string]int32
	_ = json.Unmarshal([]byte(pm.Prices), &prices)

	return &types.Module{
		Id:           pm.Id,
		SerialNumber: pm.SerialNumber,
		StationSlug:  pm.StationSlug,
		ProductSlug:  pm.ProductSlug,
		Prices:       prices,
		CreatedAt:    pm.CreatedAt,
	}
}
