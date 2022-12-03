package endpoint

import (
	"context"
	"errors"

	"github.com/khusainnov/auth-service/gen/pb"
	"github.com/khusainnov/auth-service/pkg/service"
	"google.golang.org/grpc/codes"
)

type FileService struct {
	pb.UnimplementedFileServiceServer
	services *service.Service
}

func NewWorkService(services *service.Service) *FileService {
	return &FileService{services: services}
}

func (ws *FileService) CreateFile(ctx context.Context, req *pb.UserFile) (*pb.CreateFileResponse, error) {
	logger.Printf("%+v\n", req)
	rsp, err := ws.services.CreateFile(req)
	if err != nil {
		logger.Printf("%+v\n", err)
		return nil, err
	}

	return rsp, ctx.Err()
}

func (ws *FileService) GetAllFiles(ctx context.Context, req *pb.FileRequest) (*pb.ResponseFile, error) {
	logger.Printf("%+v\n", req)
	rsp, err := ws.services.GetAllFiles(req)
	if err != nil {
		logger.Printf("%+v\n", err)
		return &pb.ResponseFile{
			Code:     int64(codes.Internal),
			Username: "",
			Chunks:   nil,
		}, errors.New(err.Error())
	}

	return rsp, ctx.Err()
}

func (ws *FileService) GetFile(ctx context.Context, req *pb.FileRequest) (*pb.ResponseFile, error) {
	logger.Printf("%+v\n", req)
	rsp, err := ws.services.GetFileByUsername(req)
	if err != nil {
		logger.Printf("%+v\n", err)
		return &pb.ResponseFile{
			Code:     int64(codes.Internal),
			Username: req.Username,
			Chunks:   nil,
		}, errors.New("files with this username doesn't exists " + err.Error())
	}

	return rsp, ctx.Err()
}

func (ws *FileService) GenerateFile(ctx context.Context, req *pb.FileRequest) (*pb.ResponseName, error) {
	logger.Printf("%+v\n", req)
	rsp, err := ws.services.GenerateFile(req)
	if err != nil {
		logger.Printf("%+v\n", err)
		return nil, err
	}

	return rsp, ctx.Err()
}
