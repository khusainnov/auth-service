package service

import (
	"github.com/khusainnov/auth-service/gen/pb"
	"github.com/khusainnov/auth-service/pkg/repository"
)

type Auth interface {
	CreateUser(u *pb.User) (*pb.ResponseMsg, error)
	GetUser(login *pb.UserRequest) (*pb.ResponseAuth, error)
	UpdateUser(u *pb.User) (*pb.ResponseMsg, error)
	DeleteUser(u *pb.DeleteRequest) (*pb.DeleteResponse, error)
	ResetPassword(login *pb.UserRequest) (*pb.ResponseMsg, error)
}

type File interface {
	CreateFile(file *pb.UserFile) (*pb.CreateFileResponse, error)
	GetFileByUsername(req *pb.FileRequest) (*pb.ResponseFile, error)
	GetAllFiles(req *pb.FileRequest) (*pb.ResponseFile, error)
	GenerateFile(req *pb.FileRequest) (*pb.ResponseName, error)
}

type Statistics interface {
	GetStatistics(req *pb.StatisticsRequest) (*pb.StatisticsResponse, error)
}

type Service struct {
	Auth
	File
	Statistics
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Auth:       NewAuthService(repo),
		File:       NewFileService(repo),
		Statistics: NewStatisticsService(repo),
	}
}
