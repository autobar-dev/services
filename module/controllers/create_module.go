package controllers

import (
	"fmt"

	"github.com/autobar-dev/services/module/types"
	"github.com/autobar-dev/services/module/utils"
)

func CreateModuleController(app_context *types.AppContext) (*types.CreateModuleResponse, error) {
	mr := app_context.Repositories.Module
	ar := app_context.Repositories.Auth

	valid_serial_number := false
	serial_number := ""

	for !valid_serial_number { // make this a forever loop?
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

	private_key, err := ar.RegisterModule(*serial_number_returned)
	if err != nil {
		fmt.Printf("IMPORTANT: failed to register module in auth service: %+v\n", err)
		return nil, err
	}

	module, err := GetModuleController(app_context, *serial_number_returned)
	if err != nil {
		return nil, err
	}

	cmr := utils.ConstructCreateModuleResponse(module, *private_key)
	return cmr, nil
}
