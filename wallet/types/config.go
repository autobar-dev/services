package types

import (
	"errors"
	"os"
	"strconv"
)

type Config struct {
	Port        int
	DatabaseURL string
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

	return &Config{
		Port:        port,
		DatabaseURL: database_url,
	}, nil
}
