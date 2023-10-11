package utils

import (
	"encoding/json"

	"github.com/autobar-dev/services/module/repositories"
	"github.com/autobar-dev/services/module/types"
)

func PostgresModuleToModule(pm repositories.PostgresModule, pdu repositories.PostgresDisplayUnit) *types.Module {
	var prices map[string]int
	_ = json.Unmarshal([]byte(pm.Prices), &prices)

	return &types.Module{
		Id:              pm.Id,
		SerialNumber:    pm.SerialNumber,
		StationId:       pm.StationId,
		ProductId:       pm.ProductId,
		Enabled:         pm.Enabled,
		Prices:          prices,
		DisplayCurrency: pm.DisplayCurrency,
		DisplayUnit: types.DisplayUnit{
			Id:                     pdu.Id,
			Amount:                 pdu.Amount,
			Symbol:                 pdu.Symbol,
			DivisorFromMillilitres: pdu.DivisorFromMillilitres,
			DecimalsDisplayed:      pdu.DecimalsDisplayed,
			CreatedAt:              pdu.CreatedAt,
			UpdatedAt:              pdu.UpdatedAt,
		},
		CreatedAt: pm.CreatedAt,
		UpdatedAt: pm.UpdatedAt,
	}
}

func RedisModuleToModule(rm repositories.RedisModule) *types.Module {
	return &types.Module{
		Id:              rm.Id,
		SerialNumber:    rm.SerialNumber,
		StationId:       rm.StationId,
		ProductId:       rm.ProductId,
		Enabled:         rm.Enabled,
		Prices:          rm.Prices,
		DisplayCurrency: rm.DisplayCurrency,
		DisplayUnit: types.DisplayUnit{
			Id:                     rm.DisplayUnit.Id,
			Symbol:                 rm.DisplayUnit.Symbol,
			DivisorFromMillilitres: rm.DisplayUnit.DivisorFromMillilitres,
			DecimalsDisplayed:      rm.DisplayUnit.DecimalsDisplayed,
			CreatedAt:              rm.DisplayUnit.CreatedAt,
			UpdatedAt:              rm.DisplayUnit.UpdatedAt,
		},
		CreatedAt: rm.CreatedAt,
		UpdatedAt: rm.UpdatedAt,
	}
}

func ModuleToRedisModule(m types.Module) *repositories.RedisModule {
	return &repositories.RedisModule{
		Id:              m.Id,
		SerialNumber:    m.SerialNumber,
		StationId:       m.StationId,
		ProductId:       m.ProductId,
		Enabled:         m.Enabled,
		Prices:          m.Prices,
		DisplayCurrency: m.DisplayCurrency,
		DisplayUnit: repositories.RedisDisplayUnit{
			Id:                     m.DisplayUnit.Id,
			Symbol:                 m.DisplayUnit.Symbol,
			DivisorFromMillilitres: m.DisplayUnit.DivisorFromMillilitres,
			DecimalsDisplayed:      m.DisplayUnit.DecimalsDisplayed,
			CreatedAt:              m.DisplayUnit.CreatedAt,
			UpdatedAt:              m.DisplayUnit.UpdatedAt,
		},
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func ConstructCreateModuleResponse(module *types.Module, private_key string) *types.CreateModuleResponse {
	return &types.CreateModuleResponse{
		Module:     module,
		PrivateKey: private_key,
	}
}
