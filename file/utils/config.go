package utils

import (
	"errors"
	"os"
	"strconv"

	"github.com/autobar-dev/services/file/types"
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

	s3_endpoint := os.Getenv("S3_ENDPOINT")
	if s3_endpoint == "" {
		return nil, errors.New("S3_ENDPOINT env var not set")
	}

	s3_access_key_id := os.Getenv("S3_ACCESS_KEY_ID")
	if s3_access_key_id == "" {
		return nil, errors.New("S3_ACCESS_KEY_ID env var not set")
	}

	s3_secret_access_key := os.Getenv("S3_SECRET_ACCESS_KEY")
	if s3_secret_access_key == "" {
		return nil, errors.New("S3_SECRET_ACCESS_KEY env var not set")
	}

	s3_bucket_name := os.Getenv("S3_BUCKET_NAME")
	if s3_bucket_name == "" {
		return nil, errors.New("S3_BUCKET_NAME env var not set")
	}

	return &types.Config{
		Port:              port,
		DatabaseURL:       database_url,
		RedisURL:          redis_url,
		S3Endpoint:        s3_endpoint,
		S3AccessKeyId:     s3_access_key_id,
		S3SecretAccessKey: s3_secret_access_key,
		S3BucketName:      s3_bucket_name,
	}, nil
}
