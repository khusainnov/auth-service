package service

import (
	"errors"
	"fmt"
	"net/smtp"
	"os"
)

var (
	host      = "smtp.gmail.com"
	port      = "587"
	address   = host + ":" + port
	tgSupport = "https://t.me/Fatik_props"
)

// TODO: function for send greeting mail after sign-up
// TODO: function for send reset password mail
// TODO: function for send news mail

func sendGreetingMail(username, email string) error {
	auth, from := authMail()

	if err := smtp.SendMail(address, auth, from, []string{email}, []byte(fmt.Sprintf("Subject: приемная комиссия БГПУ им. М. Акмуллы\n\n\n%s, благодарим за вашу регистрацию\nС уважением,\nБГПУ им. М. Акмуллы", username))); err != nil {
		return errors.New("cannot send mail, due to error: " + err.Error())
	}

	return nil
}

func sendResetPasswordMail(username, email, password string) error {
	auth, from := authMail()

	if err := smtp.SendMail(address, auth, from, []string{email}, []byte(fmt.Sprintf("Subject: Security system\n%s, ваш пароль был успешно изменен.\n\nЕсли это были не вы, пожалуйста, свяжитесь с поддержкой: %s\n\n\n\n%s, your password was successfully changed.\n\nIf it wasn't you, please contact with support: %s\n\n\n\nNew Password: %s", username, tgSupport, username, tgSupport, password))); err != nil {
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
