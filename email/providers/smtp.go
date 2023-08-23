package providers

import (
	"crypto/tls"

	"gopkg.in/gomail.v2"
)

type SmtpEmailProvider struct {
	dialer gomail.Dialer
}

func NewSmtpEmailProvider(
	hostname string,
	port int,
	username string,
	password string,
) *SmtpEmailProvider {
	d := gomail.NewDialer(hostname, port, username, password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	return &SmtpEmailProvider{dialer: *d}
}

func (sep *SmtpEmailProvider) Send(from string, to string, subject string, plain_text string, html string) error {
	m := gomail.NewMessage()

	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", plain_text)
	m.AddAlternative("text/html", html)

	return sep.dialer.DialAndSend(m)
}
