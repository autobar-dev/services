package types

import (
	"github.com/autobar-dev/services/module/repositories"
	"github.com/autobar-dev/shared-libraries/go/auth-repository"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Repositories struct {
	Module   *repositories.ModuleRepository
	Cache    *repositories.CacheRepository
	Realtime *repositories.RealtimeRepository
	Auth     *authrepository.AuthRepository
}

type AppContext struct {
	MetaFactors  *MetaFactors
	Config       *Config
	Repositories *Repositories
	AmqpChannel  *amqp.Channel
}
