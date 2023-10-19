package utils

import (
	"errors"
	"os"
	"strconv"

	"github.com/autobar-dev/services/product/types"
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

	meili_url := os.Getenv("MEILI_URL")
	if meili_url == "" {
		return nil, errors.New("MEILI_URL env var not set")
	}

	meili_api_key := os.Getenv("MEILI_API_KEY")

	file_service_url := os.Getenv("FILE_SERVICE_URL")
	if file_service_url == "" {
		return nil, errors.New("FILE_SERVICE_URL env var not set")
	}

	return &types.Config{
		Port:           port,
		DatabaseURL:    database_url,
		RedisURL:       redis_url,
		MeiliURL:       meili_url,
		MeiliApiKey:    meili_api_key,
		FileServiceURL: file_service_url,
	}, nil
}
