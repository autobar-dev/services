package types

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"go.a5r.dev/services/module/repositories"
)

type Repositories struct {
	Module   *repositories.ModuleRepository
	Auth     *repositories.AuthRepository
	Realtime *repositories.RealtimeRepository
}

type AppContext struct {
	Meta         *Meta
	Config       *Config
	Repositories *Repositories
	AmqpChannel  *amqp.Channel
}
