package controllers

import (
	"go.a5r.dev/services/wallet/types"
	"go.a5r.dev/services/wallet/utils"
)

func GetWalletController(app_context *types.AppContext, email string) (*types.Wallet, error) {
	wr := app_context.Repositories.Wallet

	pw, err := wr.Get(email)
	if err != nil {
		return nil, err
	}

	transactions, err := GetAllTransactionsController(app_context, email)
	if err != nil {
		return nil, err
	}

	wallet, err := utils.ConstructWallet(*pw, *transactions)

	if err != nil {
		return nil, err
	}

	return wallet, nil
}
