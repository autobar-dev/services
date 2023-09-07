package types

import (
	"errors"
	"os"
	"strconv"
)

type Config struct {
	Port                       int
	DatabaseURL                string
	RedisCacheURL              string
	RedisStateURL              string
	AuthServiceURL             string
	RealtimeServiceURL         string
	UserServiceURL             string
	WalletServiceURL           string
	CurrencyServiceURL         string
	ProductServiceURL          string
	AmqpURL                    string
	ModuleReportTimeoutSeconds int
	JwtSecret                  string
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

	redis_cache_url := os.Getenv("REDIS_CACHE_URL")
	if redis_cache_url == "" {
		return nil, errors.New("REDIS_CACHE_URL env var not set")
	}

	redis_state_url := os.Getenv("REDIS_STATE_URL")
	if redis_state_url == "" {
		return nil, errors.New("REDIS_STATE_URL env var not set")
	}

	amqp_url := os.Getenv("AMQP_URL")
	if amqp_url == "" {
		return nil, errors.New("AMQP_URL env var not set")
	}

	auth_service_url := os.Getenv("AUTH_SERVICE_URL")
	if auth_service_url == "" {
		return nil, errors.New("AUTH_SERVICE_URL env var not set")
	}

	realtime_service_url := os.Getenv("REALTIME_SERVICE_URL")
	if realtime_service_url == "" {
		return nil, errors.New("REALTIME_SERVICE_URL env var not set")
	}

	user_service_url := os.Getenv("USER_SERVICE_URL")
	if user_service_url == "" {
		return nil, errors.New("USER_SERVICE_URL env var not set")
	}

	wallet_service_url := os.Getenv("WALLET_SERVICE_URL")
	if wallet_service_url == "" {
		return nil, errors.New("WALLET_SERVICE_URL env var not set")
	}

	currency_service_url := os.Getenv("CURRENCY_SERVICE_URL")
	if currency_service_url == "" {
		return nil, errors.New("CURRENCY_SERVICE_URL env var not set")
	}

	product_service_url := os.Getenv("PRODUCT_SERVICE_URL")
	if product_service_url == "" {
		return nil, errors.New("PRODUCT_SERVICE_URL env var not set")
	}

	module_report_timeout_seconds := os.Getenv("MODULE_REPORT_TIMEOUT_SECONDS")
	module_report_timeout_seconds_int, err := strconv.Atoi(module_report_timeout_seconds)
	if module_report_timeout_seconds == "" || err != nil {
		return nil, errors.New("MODULE_REPORT_TIMEOUT_SECONDS env var not set")
	}

	jwt_secret := os.Getenv("JWT_SECRET")
	if jwt_secret == "" {
		return nil, errors.New("JWT_SECRET env var not set")
	}

	return &Config{
		Port:                       port,
		DatabaseURL:                database_url,
		RedisCacheURL:              redis_cache_url,
		RedisStateURL:              redis_state_url,
		AuthServiceURL:             auth_service_url,
		RealtimeServiceURL:         realtime_service_url,
		UserServiceURL:             user_service_url,
		WalletServiceURL:           wallet_service_url,
		CurrencyServiceURL:         currency_service_url,
		ProductServiceURL:          product_service_url,
		AmqpURL:                    amqp_url,
		ModuleReportTimeoutSeconds: module_report_timeout_seconds_int,
		JwtSecret:                  jwt_secret,
	}, nil
}
