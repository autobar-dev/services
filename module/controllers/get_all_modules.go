package controllers

import (
	"github.com/autobar-dev/services/module/types"
	"github.com/autobar-dev/services/module/utils"
)

func GetAllModulesController(app_context *types.AppContext) (*[]types.Module, error) {
	mr := app_context.Repositories.Module

	pms, err := mr.GetAll()
	if err != nil {
		return nil, err
	}

	modules := []types.Module{}

	for _, pm := range *pms {
		modules = append(modules, *utils.PostgresModuleToModule(pm))
	}

	return &modules, nil
}
