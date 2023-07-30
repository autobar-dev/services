package controllers

import (
	"fmt"

	"go.a5r.dev/services/module/repositories"
	"go.a5r.dev/services/module/types"
	"go.a5r.dev/services/module/utils"
)

func GetAllModulesForStationController(app_context *types.AppContext, station_id string) (*[]types.Module, error) {
	mr := app_context.Repositories.Module
	cr := app_context.Repositories.Cache

	rms, err := cr.GetAllModulesForStation(station_id)
	if err == nil {
		modules := []types.Module{}
		for _, rm := range *rms {
			modules = append(modules, *utils.RedisModuleToModule(rm))
		}

		return &modules, nil
	}

	pms, err := mr.GetAllForStation(station_id)
	if err != nil {
		return nil, err
	}

	modules := []types.Module{}
	rms_to_cache := []repositories.RedisModule{}

	for _, pm := range *pms {
		m := utils.PostgresModuleToModule(pm)

		modules = append(modules, *m)
		rms_to_cache = append(rms_to_cache, *utils.ModuleToRedisModule(*m))
	}

	err = cr.SetAllModulesForStation(station_id, rms_to_cache)
	if err != nil {
		fmt.Printf("failed to update cache when setting all modules: %v\n", err)
	}

	return &modules, nil
}
