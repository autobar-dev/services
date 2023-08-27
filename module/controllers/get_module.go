package controllers

import (
	"fmt"

	"github.com/autobar-dev/services/module/types"
	"github.com/autobar-dev/services/module/utils"
)

func GetModuleController(app_context *types.AppContext, serial_number string) (*types.Module, error) {
	mr := app_context.Repositories.Module
	cr := app_context.Repositories.Cache

	rm, err := cr.GetModule(serial_number)
	if err == nil {
		return utils.RedisModuleToModule(*rm), nil
	}

	pm, err := mr.Get(serial_number)
	if err != nil {
		return nil, err
	}

	m := utils.PostgresModuleToModule(*pm)

	err = cr.SetModule(m.Id, m.SerialNumber, m.StationId, m.ProductId, m.Enabled, m.Prices, m.CreatedAt, m.UpdatedAt)
	if err != nil {
		fmt.Printf("failed to set cache for module: %v\n", err)
	}

	return m, nil
}
