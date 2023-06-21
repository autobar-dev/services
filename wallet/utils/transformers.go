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
	}

	return tt
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

func ConstructBalanceAndCurrencyCodeFromTransactions(ts []types.Transaction, cc string) (int, string) {
	balance := 0
	currency_code := cc

	for _, transaction := range ts {
		switch transaction.TransactionType {
		case types.TransactionTypeDeposit:
			balance += transaction.Value
		case types.TransactionTypeWithdraw:
			balance -= transaction.Value
		case types.TransactionTypePurchase:
			balance -= transaction.Value
		case types.TransactionTypeRefund:
			balance += transaction.Value
		}
	}

	return balance, currency_code
}

func ConstructWallet(pw repositories.PostgresWallet, ts []types.Transaction) *types.Wallet {
	balance, currency_code := ConstructBalanceAndCurrencyCodeFromTransactions(ts, pw.CurrencyCode)

	return &types.Wallet{
		Id:           pw.Id,
		UserEmail:    pw.UserEmail,
		CurrencyCode: currency_code,
		Balance:      balance,
	}
}
