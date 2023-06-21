package controllers

import (
	"go.a5r.dev/services/wallet/types"
	"go.a5r.dev/services/wallet/utils"
)

func GetAllTransactionsController(app_context *types.AppContext, email string) (*[]types.Transaction, error) {
	wr := app_context.Repositories.Wallet
	tr := app_context.Repositories.Transaction

	pw, err := wr.Get(email)
	if err != nil {
		return nil, err
	}

	pts, err := tr.GetForWallet(pw.Id)
	if err != nil {
		return nil, err
	}

	transactions := []types.Transaction{}
	for _, pt := range *pts {
		transactions = append(transactions, *utils.PostgresTransactionToTransaction(pt))
	}

	return &transactions, nil
}
