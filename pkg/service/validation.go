package service

import (
	"net/mail"
	"regexp"
)

func validateEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return false
	}

	return true
}

func validatePhone(phone string) bool {
	//query := `^((8|\+7)[\- ]?)?(\(?\d{3}\)?[\- ]?)?[\d\- ]{7,10}$`
	val, err := regexp.Match(`^((8|\+7)[\- ]?)?(\(?\d{3}\)?[\- ]?)?[\d\- ]{7,10}$`, []byte(phone))
	if !val || err != nil {
		return false
	}

	return true
}
