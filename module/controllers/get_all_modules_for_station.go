package controllers

import (
	"fmt"

	"github.com/autobar-dev/services/module/repositories"
	"github.com/autobar-dev/services/module/types"
	"github.com/autobar-dev/services/module/utils"
)

func GetAllModulesForStationController(app_context *types.AppContext, station_id string) (*[]types.Module, error) {
	mr := app_context.Repositories.Module
	dur := app_context.Repositories.DisplayUnit
	cr := app_context.Repositories.Cache
	cur := app_context.Repositories.Currency

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
		pdu, err := dur.GetDisplayUnit(pm.DisplayUnitId)
		if err != nil {
			return nil, err
		}

		c, err := cur.GetCurrencyByCode(pm.DisplayCurrency)
		if err != nil {
			return nil, err
		}

		m := utils.PostgresModuleToModule(pm, *c, *pdu)

		modules = append(modules, *m)
		rms_to_cache = append(rms_to_cache, *utils.ModuleToRedisModule(*m))
	}

	err = cr.SetAllModulesForStation(station_id, rms_to_cache)
	if err != nil {
		fmt.Printf("failed to update cache when setting all modules: %v\n", err)
	}

	return &modules, nil
}
