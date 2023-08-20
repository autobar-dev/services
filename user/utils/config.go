package utils

import (
	"errors"
	"os"
	"strconv"

	"github.com/autobar-dev/services/user/types"
)

func LoadEnvVars() (*types.Config, error) {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		return nil, err
	}

	database_url := os.Getenv("DATABASE_URL")
	if database_url == "" {
		return nil, errors.New("DATABASE_URL env var not set")
	}

	redis_url := os.Getenv("REDIS_URL")
	if redis_url == "" {
		return nil, errors.New("REDIS_URL env var not set")
	}

	auth_service_url := os.Getenv("AUTH_SERVICE_URL")
	if auth_service_url == ""{
		return nil, errors.New("AUTH_SERVICE_URL env var not set")
	}

	return &types.Config{
		Port:        port,
		DatabaseURL: database_url,
		RedisURL:    redis_url,
		AuthServiceURL: auth_service_url,
	}, nil
}
