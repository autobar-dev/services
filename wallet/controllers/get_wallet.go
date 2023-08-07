package controllers

import (
	"fmt"

	"go.a5r.dev/services/wallet/types"
	"go.a5r.dev/services/wallet/utils"
)

func GetWalletController(app_context *types.AppContext, user_id string) (*types.Wallet, error) {
	wr := app_context.Repositories.Wallet
	car := app_context.Repositories.Cache

	cached_wallet, _ := car.GetWallet(user_id)
	if cached_wallet != nil {
		return utils.RedisWalletToWallet(*cached_wallet, user_id), nil
	}

	fmt.Println("cache miss. will update redis")

	pw, err := wr.Get(user_id)
	if err != nil {
		return nil, err
	}

	transactions, err := GetAllTransactionsController(app_context, user_id)
	if err != nil {
		return nil, err
	}

	wallet, err := utils.ConstructWallet(*pw, *transactions)
	if err != nil {
		return nil, err
	}

	if err = car.SetWallet(wallet.Id, wallet.UserId, wallet.Balance, wallet.CurrencyCode); err != nil {
		fmt.Println("failed to update cache")
	}

	return wallet, nil
}
