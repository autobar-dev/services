package utils

import (
	"errors"
	"os"
	"strconv"

	"github.com/autobar-dev/services/email/types"
)

func LoadEnvVars() (*types.Config, error) {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		return nil, err
	}

	logger_environment := os.Getenv("LOGGER_ENVIRONMENT")
	if logger_environment == "" {
		return nil, errors.New("LOGGER_ENVIRONMENT env var not set")
	}

	jwt_key := os.Getenv("JWT_KEY")
	if jwt_key == "" {
		return nil, errors.New("JWT_KEY env var not set")
	}

	smtp_hostname := os.Getenv("SMTP_HOSTNAME")
	if smtp_hostname == "" {
		return nil, errors.New("SMTP_HOSTNAME env var not set")
	}

	smtp_port, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	if err != nil {
		return nil, err
	}

	smtp_username := os.Getenv("SMTP_USERNAME")

	smtp_password := os.Getenv("SMTP_PASSWORD")

	return &types.Config{
		Port:              port,
		LoggerEnvironment: logger_environment,
		JwtKey:            jwt_key,
		SmtpHostname:      smtp_hostname,
		SmtpPort:          smtp_port,
		SmtpUsername:      smtp_username,
		SmtpPassword:      smtp_password,
	}, nil
}
