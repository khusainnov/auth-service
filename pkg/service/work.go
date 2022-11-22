package service

import (
	"bytes"
	"errors"
	"io/ioutil"
	"os"

	"github.com/google/uuid"
	"github.com/khusainnov/auth-service/gen/pb"
	"github.com/khusainnov/auth-service/pkg/repository"
	"google.golang.org/grpc/codes"
)

type WorkService struct {
	repo repository.Work
}

func NewWorkService(repo repository.Work) *WorkService {
	return &WorkService{repo: repo}
}

func (ws *WorkService) CreateFile(file *pb.UserWork) (*pb.ResponseWork, error) {
	var bb bytes.Buffer
	bb.Write(file.GetFile())

	ud, _ := uuid.NewRandom()

	f, _ := os.Create("pdf/" + ud.String() + ".pdf")

	_, err := bb.WriteTo(f)

	if err != nil {
		return &pb.ResponseWork{Code: int64(codes.Internal)}, errors.New("cannot write bytes into pdf file")
	}

	rspBytes, err := ioutil.ReadFile("pdf/" + ud.String() + ".pdf")
	if err != nil {
		return &pb.ResponseWork{Code: int64(codes.Internal)}, errors.New("cannot read bytes from pdf file")
	}

	return &pb.ResponseWork{Code: int64(codes.OK), File: rspBytes}, nil
}
