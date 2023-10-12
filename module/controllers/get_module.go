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
	cur := app_context.Repositories.Currency

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

	c, err := cur.GetCurrencyByCode(pm.DisplayCurrency)
	if err != nil {
		return nil, err
	}

	m := utils.PostgresModuleToModule(*pm, *c, *pdu)
	rm = utils.ModuleToRedisModule(*m)

	err = cr.SetModule(*rm)
	if err != nil {
		fmt.Printf("failed to set cache for module: %v\n", err)
	}

	return m, nil
}
