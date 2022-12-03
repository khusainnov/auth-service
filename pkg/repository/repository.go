package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/khusainnov/auth-service/gen/pb"
)

type Auth interface {
	CreateUser(u *pb.User) (string, error)
	GetUser(login *pb.UserRequest) (string, error)
	UpdateUser(u *pb.User) (string, error)
	DeleteUser(username string) (string, error)
	ResetPassword(username, newPassword string) (string, error)
}

type File interface {
	CreateFile(username string, fileChunksL []byte) (*pb.CreateFileResponse, error)
	GetFileByUsername(username string) (*pb.ResponseFile, error)
	GetAllFiles(username string) (*pb.ResponseFile, error)
	GenerateFile(username string) (*pb.ResponseName, error)
}

type Statistics interface {
	GetStatistics(username string) (*pb.StatisticsResponse, error)
}

type Repository struct {
	Auth
	File
	Statistics
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Auth:       NewAuthPostgres(db),
		File:       NewFilePostgres(db),
		Statistics: NewStatisticsPostgres(db),
	}
}
