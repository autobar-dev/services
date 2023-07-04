package types

import (
	sse "github.com/r3labs/sse/v2"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.a5r.dev/services/realtime/repositories"
)

type Repositories struct {
	Auth *repositories.AuthRepository
}

type AppContext struct {
	Meta         *Meta
	Config       *Config
	Repositories *Repositories
	AmqpChannel  *amqp.Channel
	SseServer    *sse.Server
}
