package types

type EmailProvider interface {
	Send(from string, to string, subject string, plain_text string, html string) error
}
