package types

import (
	"errors"
	"os"
	"strconv"
)

type Config struct {
	Port               int
	DatabaseURL        string
	AuthServiceURL     string
	CurrencyServiceURL string
	RedisURL           string
	JwtSecret          string
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

	auth_service_url := os.Getenv("AUTH_SERVICE_URL")
	if auth_service_url == "" {
		return nil, errors.New("AUTH_SERVICE_URL env var not set")
	}

	currency_service_url := os.Getenv("CURRENCY_SERVICE_URL")
	if currency_service_url == "" {
		return nil, errors.New("CURRENCY_SERVICE_URL env var not set")
	}

	redis_url := os.Getenv("REDIS_URL")
	if redis_url == "" {
		return nil, errors.New("REDIS_URL env var not set")
	}

	jwt_secret := os.Getenv("JWT_SECRET")
	if jwt_secret == "" {
		return nil, errors.New("JWT_SECRE env var not set")
	}

	return &Config{
		Port:               port,
		DatabaseURL:        database_url,
		AuthServiceURL:     auth_service_url,
		CurrencyServiceURL: currency_service_url,
		RedisURL:           redis_url,
		JwtSecret:          jwt_secret,
	}, nil
}
