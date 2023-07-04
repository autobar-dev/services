package types

import (
	"errors"
	"os"
	"strconv"
)

type Config struct {
	Port                 int
	SseHeartbeatInterval int
	AuthServiceURL       string
	AmqpURL              string
}

func LoadEnvVars() (*Config, error) {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		return nil, err
	}

	sse_heartbeat_interval, err := strconv.Atoi(os.Getenv("SSE_HEARTBEAT_INTERVAL"))
	if err != nil {
		return nil, err
	}

	amqp_url := os.Getenv("AMQP_URL")
	if amqp_url == "" {
		return nil, errors.New("AMQP_URL env var not set")
	}

	auth_service_url := os.Getenv("AUTH_SERVICE_URL")
	if auth_service_url == "" {
		return nil, errors.New("AUTH_SERVICE_URL env var not set")
	}

	return &Config{
		Port:                 port,
		SseHeartbeatInterval: sse_heartbeat_interval,
		AuthServiceURL:       auth_service_url,
		AmqpURL:              amqp_url,
	}, nil
}
