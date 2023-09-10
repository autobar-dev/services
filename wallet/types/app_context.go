package types

import (
	"github.com/autobar-dev/services/wallet/repositories"
	authrepository "github.com/autobar-dev/shared-libraries/go/auth-repository"
)

type Repositories struct {
	Auth        *authrepository.AuthRepository
	Wallet      *repositories.WalletRepository
	Transaction *repositories.TransactionRepository
	Currency    *repositories.CurrencyRepository
	Cache       *repositories.CacheRepository
}

type AppContext struct {
	MetaFactors  *MetaFactors
	Config       *Config
	Repositories *Repositories
}
