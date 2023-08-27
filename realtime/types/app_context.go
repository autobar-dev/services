package types

import (
	sse "github.com/r3labs/sse/v2"
	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/autobar-dev/services/realtime/repositories"
	"github.com/autobar-dev/shared-libraries/go/auth-repository"
)

type Repositories struct {
	Auth  *authrepository.AuthRepository
	Redis *repositories.RedisRepository
	Mq    *repositories.MqRepository
}

type AppContext struct {
	MetaFactors  *MetaFactors
	Config       *Config
	Repositories *Repositories
	AmqpChannel  *amqp.Channel
	SseServer    *sse.Server
}
