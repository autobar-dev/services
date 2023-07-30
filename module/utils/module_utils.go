package utils

import (
	"encoding/json"

	"go.a5r.dev/services/module/repositories"
	"go.a5r.dev/services/module/types"
)

func PostgresModuleToModule(pm repositories.PostgresModule) *types.Module {
	var prices map[string]int
	_ = json.Unmarshal([]byte(pm.Prices), &prices)

	return &types.Module{
		Id:           pm.Id,
		SerialNumber: pm.SerialNumber,
		StationId:    pm.StationId,
		ProductId:    pm.ProductId,
		Enabled:      pm.Enabled,
		Prices:       prices,
		CreatedAt:    pm.CreatedAt,
		UpdatedAt:    pm.UpdatedAt,
	}
}

func ConstructCreateModuleResponse(sm *repositories.ServiceModule, module *types.Module) *types.CreateModuleResponse {
	return &types.CreateModuleResponse{
		Module:     module,
		PrivateKey: sm.PrivateKey,
	}
}
