package utils

import (
	"errors"
	"os"
	"strconv"

	"github.com/autobar-dev/services/auth/types"
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

	logger_environment := os.Getenv("LOGGER_ENVIRONMENT")
	if logger_environment == "" {
		return nil, errors.New("LOGGER_ENVIRONMENT env var not set")
	}

	jwt_key := os.Getenv("JWT_KEY")
	if jwt_key == "" {
		return nil, errors.New("JWT_KEY env var not set")
	}

	return &types.Config{
		Port:              port,
		DatabaseURL:       database_url,
		LoggerEnvironment: logger_environment,
		JwtKey:            jwt_key,
	}, nil
}
