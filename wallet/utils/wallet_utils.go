package utils

import (
	"github.com/autobar-dev/services/wallet/repositories"
	"github.com/autobar-dev/services/wallet/types"
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
