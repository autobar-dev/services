package types

type Config struct {
	Port                    int
	JwtSecret               string
	DatabaseURL             string
	RedisURL                string
	AuthServiceURL          string
	EmailTemplateServiceURL string
	EmailServiceURL         string
}
