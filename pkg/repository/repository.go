package repository

import (
	"github.com/go-redis/redis/v9"
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

func NewRepository(db *sqlx.DB, rdb *redis.Client) *Repository {
	return &Repository{
		Auth: NewAuthPostgres(db),
	}
}
