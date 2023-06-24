package utils

import (
	"go.a5r.dev/services/wallet/repositories"
	"go.a5r.dev/services/wallet/types"
)

func ConstructWallet(pw repositories.PostgresWallet, ts []types.Transaction) (*types.Wallet, error) {
	wb := NewWalletBuilder()

	wb.SetId(pw.Id)
	wb.SetUserEmail(pw.UserEmail)

	for _, t := range ts {
		wb.AddTransaction(t)
	}

	return wb.Build()
}

func RedisWalletToWallet(rw repositories.RedisWallet, user_email string) *types.Wallet {
	return &types.Wallet{
		Id:           rw.Id,
		UserEmail:    user_email,
		Balance:      rw.Balance,
		CurrencyCode: rw.CurrencyCode,
	}
}
