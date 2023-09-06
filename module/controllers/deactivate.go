package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/autobar-dev/services/module/repositories"
	"github.com/autobar-dev/services/module/types"
	"github.com/autobar-dev/services/module/utils"
)

func DeactivateController(
	app_context *types.AppContext,
	user_id string,
) error {
	rr := app_context.Repositories.Realtime
	sr := app_context.Repositories.State
	wr := app_context.Repositories.Wallet

	as_id, err := sr.GetActivationSessionIdForUser(user_id)
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

	args_json_bytes, _ := json.Marshal(&args)
	args_json := string(args_json_bytes)

	err = rr.SendCommand(
		as.SerialNumber,
		repositories.ModuleServiceRealtimeClientType,
		types.DeactivateCommandName,
		args_json,
	)
	if err != nil {
		return err
	}

	transaction, err := wr.CreateTransactionPurchase(as.UserId, int64(as.Price))
	if err != nil {
		fmt.Printf("IMPORTANT: Error creating transaction: %s\n", err.Error())
	} else {
		fmt.Printf("Created transaction: %s\n", transaction.Id)
	}

	err = sr.ClearActivationSessionIdForModule(as.SerialNumber)
	if err != nil {
		fmt.Printf("IMPORTANT: Error clearing activation session ID for module: %s\n", err.Error())
	}

	err = sr.ClearActivationSessionIdForUser(user_id)
	if err != nil {
		fmt.Printf("IMPORTANT: Error clearing activation session ID for user: %s\n", err.Error())
	}

	err = sr.ClearActivationSession(as.Id)
	if err != nil {
		fmt.Printf("IMPORTANT: Error clearing activation session: %s\n", err.Error())
	}

	err = sr.ClearOtkForModule(as.SerialNumber)
	if err != nil {
		fmt.Printf("IMPORTANT: Error clearing otk for module: %s\n", err.Error())
	}

	return nil
}
