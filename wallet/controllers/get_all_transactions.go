package controllers

import (
	"github.com/autobar-dev/services/wallet/types"
	"github.com/autobar-dev/services/wallet/utils"
)

func GetAllTransactionsController(app_context *types.AppContext, user_id string) (*[]types.Transaction, error) {
	wr := app_context.Repositories.Wallet
	tr := app_context.Repositories.Transaction

	pw, err := wr.Get(user_id)
	if err != nil {
		return nil, err
	}

	pts, err := tr.GetAllForWallet(pw.Id)
	if err != nil {
		return nil, err
	}

	transactions := []types.Transaction{}
	for _, pt := range *pts {
		transactions = append(transactions, *utils.PostgresTransactionToTransaction(pt))
	}

	return &transactions, nil
}
