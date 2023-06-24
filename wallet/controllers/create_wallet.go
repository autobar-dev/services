package controllers

import (
	"errors"
	"strings"

	"go.a5r.dev/services/wallet/repositories"
	"go.a5r.dev/services/wallet/types"
)

func CreateWalletController(app_context *types.AppContext, user_email string, currency_code string) (*types.Wallet, error) {
	wr := app_context.Repositories.Wallet
	tr := app_context.Repositories.Transaction
	cr := app_context.Repositories.Currency

	currency_code = strings.ToUpper(currency_code)

	currency, err := cr.GetCurrency(currency_code)
	if err != nil {
		return nil, err
	}

	if !currency.Enabled {
		return nil, errors.New("specified currency is disabled")
	}

	pw, err := wr.Create(user_email)
	if err != nil {
		return nil, err
	}

	if _, err := tr.Create(pw.Id, repositories.PostgresTransactionTypeCurrencyChange, 0, currency_code); err != nil {
		return nil, err
	}

	return GetWalletController(app_context, user_email)
}
