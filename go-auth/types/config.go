package types

type Config struct {
	Port              int
	DatabaseURL       string
	LoggerEnvironment string
	JwtSecret         string
}
