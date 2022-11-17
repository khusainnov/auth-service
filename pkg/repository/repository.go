package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/khusainnov/auth-service/gen/pb"
)

type Auth interface {
	CreateUser(u *pb.User) (string, error)
	GetUser(login *pb.UserRequest) (*pb.User, error)
	UpdateUser(u *pb.User) (*pb.ResponseMsg, error)
}

type Repository struct {
	Auth
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Auth: NewAuthPostgres(db),
	}
}
