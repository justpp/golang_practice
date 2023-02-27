package email

import (
	"crypto/tls"
	"gopkg.in/gomail.v2"
)

type SMTPInfo struct {
	Host     string
	Port     int
	IsSSL    bool
	Username string
	Password string
	From     string
}

type Email struct {
	*SMTPInfo
}

func NewEmail(info *SMTPInfo) *Email {
	return &Email{}
}

func (e *Email) SendEmail(to []string, subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("form", e.From)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	dialer := gomail.NewDialer(e.Host, e.Port, e.Username, e.Password)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: e.IsSSL}
	return dialer.DialAndSend(m)
}
