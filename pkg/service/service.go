package service

import (
	"github.com/khusainnov/auth-service/gen/pb"
	"github.com/khusainnov/auth-service/pkg/repository"
)

type Auth interface {
	CreateUser(u *pb.User) (*pb.ResponseMsg, error)
	GetUser(login *pb.UserRequest) (*pb.ResponseToken, error)
	UpdateUser(u *pb.User) (*pb.ResponseMsg, error)
}

type Work interface {
	CreateFile(file *pb.UserWork) (*pb.ResponseWork, error)
}

type Service struct {
	Auth
	Work
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Auth: NewAuthService(repo),
		Work: NewWorkService(repo),
	}
}
