package types

type Config struct {
	Port                    int
	JwtSecret               string
	DatabaseURL             string
	RedisURL                string
	AuthServiceURL          string
	WalletServiceURL        string
	EmailTemplateServiceURL string
	EmailServiceURL         string
}
