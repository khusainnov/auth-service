package service

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	tokenSalt = os.Getenv("TOKEN_SALT")
	timeTS    = time.Now().Unix()
	timeTL    = time.Now().Add(time.Hour * 12).Unix()
)

type tokenClaims struct {
	jwt.StandardClaims
	Username string `json:"username"`
}

func GenerateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  timeTS,
			ExpiresAt: timeTL,
		},
		Username: username,
	})

	return token.SignedString([]byte(tokenSalt))
}

// +447539009672
func validateToken(token string) (bool, string) {
	t, err := jwt.ParseWithClaims(token, &tokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("TOKEN_SALT")), nil
	})

	if err != nil {
		return false, ""
	}

	claims := t.Claims.(*tokenClaims)
	return t.Valid, claims.Username
}
