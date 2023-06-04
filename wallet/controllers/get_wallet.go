package controllers

import "go.a5r.dev/services/wallet/types"

func GetWalletController(app_context *types.AppContext) (*types.Wallet, error) {
	wr = app_context.Repositories.Wallet

}
