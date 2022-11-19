package service

import (
	"net/smtp"
	"os"
)

var (
	from     = os.Getenv("MAIL_FROM")
	password = os.Getenv("MAIL_PASSWORD")
	host     = "smtp.gmail.com"
	port     = "587"
	address  = host + ":" + port
)

func sendMail(msg, email string) error {
	auth := smtp.PlainAuth(
		"",
		from,
		password,
		host,
	)

	if err := smtp.SendMail(address, auth, from, []string{email}, []byte("Subject: Your token\nHello little pussy with token:\n"+msg)); err != nil {
		return err
	}

	return nil
}
