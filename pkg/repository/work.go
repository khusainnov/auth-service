package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/khusainnov/auth-service/gen/pb"
)

type WorkPostgres struct {
	db *sqlx.DB
}

func NewWorkPostgres(db *sqlx.DB) *WorkPostgres {
	return &WorkPostgres{db: db}
}

func (wp *WorkPostgres) CreateFile(username, fileURL string) (*pb.ResponseWork, error) {

	return nil, nil
}
