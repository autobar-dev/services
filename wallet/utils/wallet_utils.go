package utils

import (
	"go.a5r.dev/services/wallet/repositories"
	"go.a5r.dev/services/wallet/types"
)

func ConstructWallet(pw repositories.PostgresWallet, ts []types.Transaction) (*types.Wallet, error) {
	wb := NewWalletBuilder()

	wb.SetId(pw.Id)
	wb.SetUserId(pw.UserId)

	for _, t := range ts {
		wb.AddTransaction(t)
	}

	return wb.Build()
}

func RedisWalletToWallet(rw repositories.RedisWallet, user_id string) *types.Wallet {
	return &types.Wallet{
		Id:           rw.Id,
		UserId:       user_id,
		Balance:      rw.Balance,
		CurrencyCode: rw.CurrencyCode,
	}
}
