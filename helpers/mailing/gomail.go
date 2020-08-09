package mailing

import (
	"gopkg.in/gomail.v2"
	"os"
)

type GoMail struct {
	Dialer *gomail.Dialer
}

func(mail GoMail) SendEmail(to, subject, message string) (err error){
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", os.Getenv("MAIL_SENDER"))
	mailer.SetHeader("To", to)
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/html", message)

	err = mail.Dialer.DialAndSend(mailer)
	if err != nil {
		return err
	}

	return nil
}