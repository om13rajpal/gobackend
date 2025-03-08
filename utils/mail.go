package utils

import (
	"github.com/om13rajpal/gobackend/config"
	"gopkg.in/gomail.v2"
)

func SendMail(to string, subject string, body string) error {
	mailer := gomail.NewDialer("smtp.gmail.com", 587, config.EMAIL, config.PASSWORD)

	m := gomail.NewMessage()

	m.SetHeader("Subject", subject)
	m.SetHeader("To", to)
	m.SetHeader("From", config.EMAIL)
	m.SetBody("text/html", body)

	err := mailer.DialAndSend(m)

	if err != nil {
		return err
	}

	return nil
}