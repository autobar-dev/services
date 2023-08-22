package types

type Config struct {
	Port              int
	DatabaseURL       string
	LoggerEnvironment string
	JwtKey            string
}
