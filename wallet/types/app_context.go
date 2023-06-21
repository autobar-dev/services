package types

import (
	"github.com/jmoiron/sqlx"
	"go.a5r.dev/services/wallet/repositories"
)

type Repositories struct {
	Wallet      *repositories.WalletRepository
	Transaction *repositories.TransactionRepository
}

type AppContext struct {
	Message      string
	Database     *sqlx.DB
	Repositories *Repositories
}
