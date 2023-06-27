package types

import (
	"errors"
	"os"
	"strconv"
)

type Config struct {
	Port                       int
	DatabaseURL                string
	AuthServiceURL             string
	RealtimeServiceURL         string
	AmqpURL                    string
	ModuleReportTimeoutSeconds int
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

	module_report_timeout_seconds := os.Getenv("MODULE_REPORT_TIMEOUT_SECONDS")
	module_report_timeout_seconds_int, err := strconv.Atoi(module_report_timeout_seconds)
	if module_report_timeout_seconds == "" || err != nil {
		return nil, errors.New("MODULE_REPORT_TIMEOUT_SECONDS env var not set")
	}

	return &Config{
		Port:                       port,
		DatabaseURL:                database_url,
		AuthServiceURL:             auth_service_url,
		RealtimeServiceURL:         realtime_service_url,
		AmqpURL:                    amqp_url,
		ModuleReportTimeoutSeconds: module_report_timeout_seconds_int,
	}, nil
}
