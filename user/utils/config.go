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

	jwt_secret := os.Getenv("JWT_SECRET")
	if jwt_secret == "" {
		return nil, errors.New("JWT_SECRET env var not set")
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
	if auth_service_url == "" {
		return nil, errors.New("AUTH_SERVICE_URL env var not set")
	}

	wallet_service_url := os.Getenv("WALLET_SERVICE_URL")
	if wallet_service_url == "" {
		return nil, errors.New("WALLET_SERVICE_URL env var not set")
	}

	emailtemplate_service_url := os.Getenv("EMAILTEMPLATE_SERVICE_URL")
	if emailtemplate_service_url == "" {
		return nil, errors.New("EMAILTEMPLATE_SERVICE_URL env var not set")
	}

	email_service_url := os.Getenv("EMAIL_SERVICE_URL")
	if email_service_url == "" {
		return nil, errors.New("EMAIL_SERVICE_URL env var not set")
	}

	return &types.Config{
		Port:                    port,
		JwtSecret:               jwt_secret,
		DatabaseURL:             database_url,
		RedisURL:                redis_url,
		AuthServiceURL:          auth_service_url,
		WalletServiceURL:        wallet_service_url,
		EmailTemplateServiceURL: emailtemplate_service_url,
		EmailServiceURL:         email_service_url,
	}, nil
}
