package controllers

import (
	"go.a5r.dev/services/module/types"
	"go.a5r.dev/services/module/utils"
)

func GetModuleController(app_context *types.AppContext, serial_number string) (*types.Module, error) {
	mr := app_context.Repositories.Module

	pm, err := mr.Get(serial_number)
	if err != nil {
		return nil, err
	}

	module := utils.PostgresModuleToModule(*pm)

	return module, nil
}
