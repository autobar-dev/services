package controllers

import (
	"fmt"

	"github.com/autobar-dev/services/module/repositories"
	"github.com/autobar-dev/services/module/types"
	"github.com/autobar-dev/services/module/utils"
)

func DeactivateAsModuleController(
	app_context *types.AppContext,
	serial_number string,
) error {
	rr := app_context.Repositories.Realtime
	sr := app_context.Repositories.State
	wr := app_context.Repositories.Wallet

	as_id, err := sr.GetActivationSessionIdForModule(serial_number)
	if err != nil {
		return err
	}

	ras, err := sr.GetActivationSession(*as_id)
	if err != nil {
		fmt.Printf("IMPORTANT: Error getting activation session when ID exists: %s\n", err.Error())
		return err
	}

	as := utils.RedisActivationSessionToActivationSession(*ras)

	args := &types.DeactivateCommandArgs{}

	args_map, err := utils.StructToJsonMap(args)
	if err != nil {
		return err
	}

	if as.AmountMillilitres > 0 {
		transaction, err := wr.CreateTransactionPurchase(as.UserId, int64(as.Price))
		if err != nil {
			fmt.Printf("IMPORTANT: Error creating transaction: %s\n", err.Error())
		} else {
			fmt.Printf("Created transaction: %s\n", transaction.Id)
		}
	} else {
		fmt.Println("Skipping creating transaction since amount millilitres is 0")
	}

	err = sr.ClearOtkForModule(as.SerialNumber)
	if err != nil {
		fmt.Printf("IMPORTANT: Error clearing otk for module: %s\n", err.Error())
	}

	err = sr.ClearActivationSessionIdForModule(as.SerialNumber)
	if err != nil {
		fmt.Printf("IMPORTANT: Error clearing activation session ID for module: %s\n", err.Error())
	}

	err = sr.ClearActivationSessionIdForUser(as.UserId)
	if err != nil {
		fmt.Printf("IMPORTANT: Error clearing activation session ID for user: %s\n", err.Error())
	}

	err = sr.ClearActivationSession(as.Id)
	if err != nil {
		fmt.Printf("IMPORTANT: Error clearing activation session: %s\n", err.Error())
	}

	err = rr.SendCommand(
		as.SerialNumber,
		repositories.ModuleServiceRealtimeClientType,
		types.DeactivateCommandName,
		args_map,
	)
	if err != nil {
		return err
	}

	return nil
}
