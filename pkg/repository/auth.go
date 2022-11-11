package repository

import (
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/khusainnov/auth-service/gen/pb"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (ap *AuthPostgres) CreateUser(u *pb.User) (int, error) {
	if u == nil {
		return 0, errors.New(fmt.Sprintf("Cannot create, due to empty user: \n%+v\n", u))
	}
	fmt.Printf("%+v\nCREATED\n", u)
	return 1, nil
}
