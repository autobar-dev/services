package types

import (
	"errors"
	"os"
	"strconv"
)

type Config struct {
	Port                        int
	SseHeartbeatIntervalSeconds int
	AuthServiceURL              string
	AmqpURL                     string
	RedisURL                    string
	ServiceBasepath             string
	JwtSecret                   string
}

func LoadEnvVars() (*Config, error) {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		return nil, err
	}

	sse_heartbeat_interval, err := strconv.Atoi(os.Getenv("SSE_HEARTBEAT_INTERVAL_SECONDS"))
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

	redis_url := os.Getenv("REDIS_URL")
	if redis_url == "" {
		return nil, errors.New("REDIS_URL env var not set")
	}

	service_basepath := os.Getenv("SERVICE_BASEPATH")

	jwt_secret := os.Getenv("JWT_SECRET")
	if jwt_secret == "" {
		return nil, errors.New("JWT_SECRET env var not set")
	}

	return &Config{
		Port:                        port,
		SseHeartbeatIntervalSeconds: sse_heartbeat_interval,
		AuthServiceURL:              auth_service_url,
		AmqpURL:                     amqp_url,
		RedisURL:                    redis_url,
		ServiceBasepath:             service_basepath,
		JwtSecret:                   jwt_secret,
	}, nil
}
