package controllers

import (
	"errors"
	"fmt"
	"strings"

	"go.a5r.dev/services/wallet/types"
	"go.a5r.dev/services/wallet/utils"
)

func CreateTransactionController(app_context *types.AppContext, user_email string, transaction_type types.TransactionType, value int, currency_code string) (*types.Transaction, error) {
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

	wallet, err := GetWalletController(app_context, user_email)
	if err != nil {
		return nil, err
	}

	can_perform_transaction := utils.CanPerformTransaction(wallet, transaction_type, value, currency.Code)
	if !can_perform_transaction {
		return nil, errors.New("cannot perform transaction most likely due to invalid input")
	}

	ptt := utils.TransactionTypeToPostgresTransactionType(transaction_type)

	transaction_id, err := tr.Create(wallet.Id, ptt, value, currency_code)
	if err != nil {
		return nil, err
	}

	pt, err := tr.Get(*transaction_id)
	if err != nil {
		return nil, err
	}

	transaction := utils.PostgresTransactionToTransaction(*pt)

	return transaction, nil
}
