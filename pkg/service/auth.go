package service

import (
	"crypto/sha1"
	"fmt"

	"github.com/khusainnov/auth-service/gen/pb"
	"github.com/khusainnov/auth-service/pkg/repository"
)

const (
	signingKey = "vg23k4hgk23b2jhb3j1hn"
)

type AuthService struct {
	repo repository.Auth
}

func NewAuthService(repo repository.Auth) *AuthService {
	return &AuthService{repo: repo}
}

func (as *AuthService) CreateUser(u *pb.User) (int, error) {
	u.Password = generatePasswordHash(u.Password)
	return as.repo.CreateUser(u)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%+x", hash.Sum([]byte(signingKey)))
}
