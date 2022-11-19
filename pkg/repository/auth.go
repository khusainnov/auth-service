package repository

import (
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/khusainnov/auth-service/gen/pb"
	"github.com/khusainnov/auth-service/internal/entity"
)

type AuthPostgres struct {
	db *sqlx.DB
}

const (
	userTable = "user_table"
)

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (ap *AuthPostgres) CreateUser(u *pb.User) (string, error) {
	var username string
	query := fmt.Sprintf("INSERT INTO %s VALUES ($1, $2, $3, $4, $5, $6) RETURNING username;", userTable)

	row := ap.db.QueryRow(query, u.Username, u.Name, u.Surname, u.Patronymic, u.Email, u.Password)
	if err := row.Scan(&username); err != nil {
		return "", errors.New(fmt.Sprintf("Cannot create due to error: %s", err.Error()))
	}

	return username, nil
}

func (ap *AuthPostgres) GetUser(login *pb.UserRequest) (*pb.User, error) {
	var u entity.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE (username=$1 and password_hash=$2) or (email=$1 and password_hash=$2);", userTable)

	if err := ap.db.Get(&u, query, login.GetUsername(), login.GetPassword()); err != nil {
		return nil, errors.New("incorrect input data, please check and try again")
	}

	fmt.Sprintf("\n\n%+v\n\n", u)

	return &pb.User{
		Username:   u.Username,
		Name:       u.Name,
		Surname:    u.Surname,
		Patronymic: u.Patronymic,
		Email:      u.Email,
		Password:   u.Password,
	}, nil
}

func (ap *AuthPostgres) UpdateUser(u *pb.User) (*pb.ResponseMsg, error) {
	return nil, nil
}
