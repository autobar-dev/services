package types

type Config struct {
	Port              int
	LoggerEnvironment string
	SmtpHostname      string
	SmtpPort          int
	SmtpUsername      string
	SmtpPassword      string
}
