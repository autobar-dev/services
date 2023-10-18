package types

type Config struct {
	Port              int
	DatabaseURL       string
	RedisURL          string
	S3Endpoint        string
	S3AccessKeyId     string
	S3SecretAccessKey string
	S3BucketName      string
}
