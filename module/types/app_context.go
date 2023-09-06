package types

import (
	"github.com/autobar-dev/services/module/repositories"
	"github.com/autobar-dev/shared-libraries/go/auth-repository"
	"github.com/autobar-dev/shared-libraries/go/currency-repository"
	"github.com/autobar-dev/shared-libraries/go/user-repository"
	"github.com/autobar-dev/shared-libraries/go/wallet-repository"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Repositories struct {
	Module   *repositories.ModuleRepository
	Cache    *repositories.CacheRepository
	State    *repositories.StateRepository
	Realtime *repositories.RealtimeRepository
	Auth     *authrepository.AuthRepository
	User     *userrepository.UserRepository
	Wallet   *walletrepository.WalletRepository
	Currency *currencyrepository.CurrencyRepository
}

type AppContext struct {
	MetaFactors  *MetaFactors
	Config       *Config
	Repositories *Repositories
	AmqpChannel  *amqp.Channel
}
