package types

type Config struct {
	Port           int
	DatabaseURL    string
	RedisURL       string
	MeiliURL       string
	MeiliApiKey    string
	FileServiceURL string
}
