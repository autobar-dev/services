package types

type Config struct {
	Port              int
	LoggerEnvironment string
	JwtKey            string
	SmtpHostname      string
	SmtpPort          int
	SmtpUsername      string
	SmtpPassword      string
}
