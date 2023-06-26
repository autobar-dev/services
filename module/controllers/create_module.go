package controllers

import (
	"go.a5r.dev/services/module/types"
	"go.a5r.dev/services/module/utils"
)

func CreateModuleController(app_context *types.AppContext) (*types.Module, error) {
	mr := app_context.Repositories.Module

	valid_serial_number := false
	serial_number := ""

	for !valid_serial_number {
		serial_number = utils.GenerateSerialNumber(types.SerialNumberLength)

		_, err := mr.Get(serial_number)
		if err != nil { // module with that serial number not found
			break
		}
	}

	serial_number_returned, err := mr.Create(serial_number)
	if err != nil {
		return nil, err
	}

	module, err := GetModuleController(app_context, *serial_number_returned)

	return module, nil
}
