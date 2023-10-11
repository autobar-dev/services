package controllers

import (
	"github.com/autobar-dev/services/module/types"
	"github.com/autobar-dev/services/module/utils"
)

func GetAllModulesController(app_context *types.AppContext) (*[]types.Module, error) {
	mr := app_context.Repositories.Module
	dur := app_context.Repositories.DisplayUnit

	pms, err := mr.GetAll()
	if err != nil {
		return nil, err
	}

	modules := []types.Module{}

	for _, pm := range *pms {
		pdu, err := dur.GetDisplayUnit(pm.DisplayUnitId)
		if err != nil {
			return nil, err
		}

		modules = append(modules, *utils.PostgresModuleToModule(pm, *pdu))
	}

	return &modules, nil
}
