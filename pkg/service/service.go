package service

import (
	"github.com/khusainnov/auth-service/gen/pb"
	"github.com/khusainnov/auth-service/pkg/repository"
)

type Auth interface {
	CreateUser(u *pb.User) (*pb.ResponseMsg, error)
	GetUser(login *pb.UserRequest) (*pb.User, error)
	UpdateUser(u *pb.User) (*pb.ResponseMsg, error)
}

type Service struct {
	Auth
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Auth: NewAuthService(repo),
	}
}
