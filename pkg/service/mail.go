package service

import (
	"net/smtp"
	"os"
)

var (
	host    = "smtp.gmail.com"
	port    = "587"
	address = host + ":" + port
)

// TODO: function for send greeting mail after sign-up
// TODO: function for send reset password mail
// TODO: function for send news mail

func sendMail(msg, email string) error {
	auth, from := authMail()

	if err := smtp.SendMail(address, auth, from, []string{email}, []byte("Subject: Your token\nHello little pussy with token:\n"+msg)); err != nil {
		return err
	}

	return nil
}

func resetPassword(username, password, email string) error {
	auth, from := authMail()

	if err := smtp.SendMail(address, auth, from, []string{email}, []byte("Subject: Your reset password link\nLink: localhost:8080/reset?username="+username)); err != nil {
		return err
	}

	return nil
}

func authMail() (smtp.Auth, string) {
	from := os.Getenv("MAIL_FROM")
	password := os.Getenv("MAIL_PASSWORD")
	auth := smtp.PlainAuth(
		"",
		from,
		password,
		host,
	)

	return auth, from
}
