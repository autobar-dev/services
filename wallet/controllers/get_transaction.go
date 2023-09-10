package controllers

import (
	"github.com/autobar-dev/services/wallet/types"
	"github.com/autobar-dev/services/wallet/utils"
)

func GetTransactionController(app_context *types.AppContext, id string) (*types.Transaction, error) {
	tr := app_context.Repositories.Transaction

	pt, err := tr.Get(id)
	if err != nil {
		return nil, err
	}

	return utils.PostgresTransactionToTransaction(*pt), nil
}
