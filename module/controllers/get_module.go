package controllers

import (
	"fmt"

	"github.com/autobar-dev/services/module/types"
	"github.com/autobar-dev/services/module/utils"
)

func GetModuleController(app_context *types.AppContext, serial_number string) (*types.Module, error) {
	mr := app_context.Repositories.Module
	dur := app_context.Repositories.DisplayUnit
	cr := app_context.Repositories.Cache

	rm, err := cr.GetModule(serial_number)
	if err == nil {
		return utils.RedisModuleToModule(*rm), nil
	}

	pm, err := mr.Get(serial_number)
	if err != nil {
		return nil, err
	}

	pdu, err := dur.GetDisplayUnit(pm.DisplayUnitId)
	if err != nil {
		return nil, err
	}

	m := utils.PostgresModuleToModule(*pm, *pdu)

	err = cr.SetModule(
		m.Id,
		m.SerialNumber,
		m.StationId,
		m.ProductId,
		m.Enabled,
		m.Prices,
		m.CreatedAt,
		m.UpdatedAt,
		m.DisplayUnit.Id,
		m.DisplayUnit.Symbol,
		m.DisplayUnit.DivisorFromMillilitres,
		m.DisplayUnit.DecimalsDisplayed,
		m.DisplayUnit.CreatedAt,
		m.DisplayUnit.UpdatedAt,
	)
	if err != nil {
		fmt.Printf("failed to set cache for module: %v\n", err)
	}

	return m, nil
}
