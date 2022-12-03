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

const (
	userTable       = "user_table"
	roleTable       = "role_table"
	statisticsTable = "statistics_table"
)

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (ap *AuthPostgres) CreateUser(u *pb.User) (string, error) {
	var username string
	query := fmt.Sprintf(`INSERT INTO %s (username, "name", surname, patronymic, email, password_hash, role_id) VALUES ($1, $2, $3, $4, $5, $6, 2) RETURNING username;`, userTable)

	row := ap.db.QueryRow(query, u.Username, u.Name, u.Surname, u.Patronymic, u.Email, u.Password)
	if err := row.Scan(&username); err != nil {
		return "", errors.New(fmt.Sprintf("Cannot create due to error: %s", err.Error()))
	}

	return username, nil
}

func (ap *AuthPostgres) GetUser(login *pb.UserRequest) (string, error) {
	var roleName string
	query := fmt.Sprintf("select rt.role_name from %s rt join %s ut on rt.role_id = ut.role_id WHERE (username=$1 and password_hash=$2) or (email=$1 and password_hash=$2);", roleTable, userTable)

	if err := ap.db.Get(&roleName, query, login.GetUsername(), login.GetPassword()); err != nil {
		return "", errors.New("incorrect input data, please check and try again")
	}

	fmt.Printf("\n\n%+v\n\n", roleName)

	return roleName, nil
}

func (ap *AuthPostgres) UpdateUser(u *pb.User) (string, error) {
	query := fmt.Sprintf("UPDATE %s SET name=$1, surname=$2, patronymic=$3, email=$4, password_hash=$5 WHERE username=$6;", userTable)

	if err := ap.db.QueryRow(query, u.Name, u.Surname, u.Patronymic, u.Email, u.Password, u.Username).Err(); err != nil {
		return "", errors.New(err.Error())
	}

	return "UPDATED", nil
}

func (ap *AuthPostgres) DeleteUser(username string) (string, error) {
	var userExists bool
	query := fmt.Sprintf("DELETE FROM %s WHERE username=$1;", userTable)
	queryExists := fmt.Sprintf("SELECT exists (SELECT 1 FROM %s WHERE username=$1);", userTable)
	if err := ap.db.Get(&userExists, queryExists, username); err != nil {
		return err.Error(), err
	}

	if !userExists {
		return "", errors.New("user doesn't exists")
	}

	if _, err := ap.db.Query(query, username); err != nil {
		return "", err
	}

	return "DELETED", nil
}

func (ap *AuthPostgres) ResetPassword(username, newPassword string) (string, error) {
	var email string
	query := fmt.Sprintf("UPDATE %s SET password_hash=$1 WHERE username=$2;", userTable)
	queryMail := fmt.Sprintf("SELECT email FROM %s WHERE username=$1", userTable)

	if err := ap.db.QueryRow(query, newPassword, username).Err(); err != nil {
		return "DECLINED", errors.New(err.Error())
	}

	if err := ap.db.Get(&email, queryMail, username); err != nil {
		return "cannot get EMAIL", err
	}

	return email, nil
}
