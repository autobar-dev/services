package types

import (
	"errors"
	"os"
	"strconv"
)

type Config struct {
	Port           int
	DatabaseURL    string
	AuthServiceURL string
}

func LoadEnvVars() (*Config, error) {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		return nil, err
	}

	database_url := os.Getenv("DATABASE_URL")
	if database_url == "" {
		return nil, errors.New("DATABASE_URL env var not set")
	}

	currency_service_url := os.Getenv("AUTH_SERVICE_URL")
	if currency_service_url == "" {
		return nil, errors.New("AUTH_SERVICE_URL env var not set")
	}

	return &Config{
		Port:           port,
		DatabaseURL:    database_url,
		AuthServiceURL: currency_service_url,
	}, nil
}
