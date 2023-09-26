package types

import (
	"errors"
	"os"
	"strconv"
)

type Config struct {
	Port           int
	AuthServiceURL string
	AmqpURL        string
	RedisStateURL  string
	JwtSecret      string
}

func LoadEnvVars() (*Config, error) {
	port, err := strconv.Atoi(os.Getenv("PORT"))
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

	redis_state_url := os.Getenv("REDIS_STATE_URL")
	if redis_state_url == "" {
		return nil, errors.New("REDIS_STATE_URL env var not set")
	}

	jwt_secret := os.Getenv("JWT_SECRET")
	if jwt_secret == "" {
		return nil, errors.New("JWT_SECRET env var not set")
	}

	return &Config{
		Port:           port,
		AuthServiceURL: auth_service_url,
		AmqpURL:        amqp_url,
		RedisStateURL:  redis_state_url,
		JwtSecret:      jwt_secret,
	}, nil
}
