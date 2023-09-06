package controllers

import (
	"errors"

	"github.com/autobar-dev/services/module/types"
	sharedutils "github.com/autobar-dev/shared-libraries/go/shared-utils"
)

func PrepareModuleController(app_context *types.AppContext, serial_number string) (*types.PrepareModuleData, error) {
	sr := app_context.Repositories.State

	m, err := GetModuleController(app_context, serial_number)
	if err != nil {
		return nil, err
	}

	if !m.Enabled || m.ProductId == nil {
		return nil, errors.New("module is not enabled or does not have a product id")
	}

	new_otk := sharedutils.GenerateRandomString(types.OtkLength, sharedutils.LowercaseUppercaseNumberCharacterSet)

	err = sr.SetOtkForModule(m.SerialNumber, new_otk)
	if err != nil {
		return nil, err
	}

	pmd := &types.PrepareModuleData{
		Otk:    new_otk,
		Module: *m,
	}

	return pmd, nil
}
