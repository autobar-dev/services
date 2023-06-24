package types

import (
	"github.com/jmoiron/sqlx"
	"go.a5r.dev/services/wallet/repositories"
)

type Repositories struct {
	Wallet      *repositories.WalletRepository
	Transaction *repositories.TransactionRepository
	Currency    *repositories.CurrencyRepository
}

type AppContext struct {
	Meta         *Meta
	Database     *sqlx.DB
	Repositories *Repositories
}
