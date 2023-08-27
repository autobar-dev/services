package utils

import (
	"encoding/json"

	"github.com/autobar-dev/services/module/repositories"
	"github.com/autobar-dev/services/module/types"
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

func RedisModuleToModule(rm repositories.RedisModule) *types.Module {
	return &types.Module{
		Id:           rm.Id,
		SerialNumber: rm.SerialNumber,
		StationId:    rm.StationId,
		ProductId:    rm.ProductId,
		Enabled:      rm.Enabled,
		Prices:       rm.Prices,
		CreatedAt:    rm.CreatedAt,
		UpdatedAt:    rm.UpdatedAt,
	}
}

func ModuleToRedisModule(m types.Module) *repositories.RedisModule {
	return &repositories.RedisModule{
		Id:           m.Id,
		SerialNumber: m.SerialNumber,
		StationId:    m.StationId,
		ProductId:    m.ProductId,
		Enabled:      m.Enabled,
		Prices:       m.Prices,
		CreatedAt:    m.CreatedAt,
		UpdatedAt:    m.UpdatedAt,
	}
}

func ConstructCreateModuleResponse(module *types.Module, private_key string) *types.CreateModuleResponse {
	return &types.CreateModuleResponse{
		Module:     module,
		PrivateKey: private_key,
	}
}
