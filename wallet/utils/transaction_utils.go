package utils

import "go.a5r.dev/services/wallet/types"
import "go.a5r.dev/services/wallet/repositories"

func PostgresTransactionTypeToTransactionType(ptt repositories.PostgresTransactionType) types.TransactionType {
	var tt types.TransactionType

	switch ptt {
	case repositories.PostgresTransactionTypeDeposit:
		tt = types.TransactionTypeDeposit
	case repositories.PostgresTransactionTypeWithdraw:
		tt = types.TransactionTypeWithdraw
	case repositories.PostgresTransactionTypePurchase:
		tt = types.TransactionTypePurchase
	case repositories.PostgresTransactionTypeRefund:
		tt = types.TransactionTypeRefund
	case repositories.PostgresTransactionTypeCurrencyChange:
		tt = types.TransactionTypeCurrencyChange
	}

	return tt
}
func TransactionTypeToPostgresTransactionType(tt types.TransactionType) repositories.PostgresTransactionType {
	var ptt repositories.PostgresTransactionType

	switch tt {
	case types.TransactionTypeDeposit:
		ptt = repositories.PostgresTransactionTypeDeposit
	case types.TransactionTypeWithdraw:
		ptt = repositories.PostgresTransactionTypeWithdraw
	case types.TransactionTypePurchase:
		ptt = repositories.PostgresTransactionTypePurchase
	case types.TransactionTypeRefund:
		ptt = repositories.PostgresTransactionTypeRefund
	case types.TransactionTypeCurrencyChange:
		ptt = repositories.PostgresTransactionTypeCurrencyChange
	}

	return ptt
}

func PostgresTransactionToTransaction(pt repositories.PostgresTransaction) *types.Transaction {
	return &types.Transaction{
		Id:              pt.Id,
		WalletId:        pt.WalletId,
		TransactionType: PostgresTransactionTypeToTransactionType(pt.TransactionType),
		Value:           pt.Value,
		CurrencyCode:    pt.CurrencyCode,
		CreatedAt:       pt.CreatedAt.UTC(),
	}
}

func CanPerformTransaction(wallet *types.Wallet, transaction_type types.TransactionType, value int, currency_code string) bool {
	allow_transaction := false
	switch transaction_type {
	case types.TransactionTypeDeposit:
		allow_transaction = currency_code == wallet.CurrencyCode
	case types.TransactionTypeWithdraw:
		allow_transaction = value <= wallet.Balance && currency_code == wallet.CurrencyCode
	case types.TransactionTypePurchase:
		allow_transaction = value <= wallet.Balance && currency_code == wallet.CurrencyCode
	case types.TransactionTypeRefund:
		allow_transaction = currency_code == wallet.CurrencyCode
	case types.TransactionTypeCurrencyChange:
		allow_transaction = currency_code != wallet.CurrencyCode
	}

	return allow_transaction
}
