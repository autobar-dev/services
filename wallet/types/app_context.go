package types

import (
	"go.a5r.dev/services/wallet/repositories"
)

type Repositories struct {
	Wallet      *repositories.WalletRepository
	Transaction *repositories.TransactionRepository
	Currency    *repositories.CurrencyRepository
	Cache       *repositories.CacheRepository
}

type AppContext struct {
	Meta         *Meta
	Repositories *Repositories
}
