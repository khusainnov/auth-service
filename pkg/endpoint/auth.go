package endpoint

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/khusainnov/auth-service/gen/pb"
	"github.com/khusainnov/auth-service/pkg/service"
	"github.com/khusainnov/logging"
)

var (
	logger = logging.GetLogger()
)

type AuthService struct {
	pb.UnimplementedAuthServiceServer
	sync.Mutex
	srv *service.Service
}

func NewAuthService(srv *service.Service) *AuthService {
	return &AuthService{srv: srv}
}

func (s *AuthService) CreateUser(ctx context.Context, req *pb.User) (*pb.ResponseMsg, error) {
	logger.Printf("%+v\n", req)
	s.Lock()
	rspMsg, err := s.srv.CreateUser(req)
	s.Unlock()
	if err != nil {
		logger.Printf("%+v\n", err)
		return nil, errors.New(err.Error())
	}
	return rspMsg, ctx.Err()
}

func (s *AuthService) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.ResponseAuth, error) {
	logger.Printf("%+v\n", req)
	rspUser, err := s.srv.GetUser(req)
	if err != nil {
		logger.Printf("%+v\n", err)
		return nil, errors.New(fmt.Sprintf("cannot get user due to %v", err))
	}

	return rspUser, ctx.Err()
}

func (s *AuthService) UpdateUser(ctx context.Context, req *pb.User) (*pb.ResponseMsg, error) {
	logger.Printf("%+v\n", req)
	rspMsg, err := s.srv.UpdateUser(req)
	if err != nil {
		logger.Printf("%+v\n", err)
		return nil, errors.New(fmt.Sprintf("cannot update user due to error: %v", err))
	}

	return rspMsg, ctx.Err()
}

func (s *AuthService) ResetPassword(ctx context.Context, req *pb.UserRequest) (*pb.ResponseMsg, error) {
	fmt.Printf("%+v\n\n", req)
	rsp, err := s.srv.ResetPassword(req)
	if err != nil {
		return rsp, err
	}

	return rsp, ctx.Err()
}

func (s *AuthService) DeleteUser(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	logger.Printf("%+v\n", req.Username)
	rsp, err := s.srv.DeleteUser(req)
	if err != nil {
		logger.Printf("%+v\n", err)
		return rsp, err
	}
	return rsp, ctx.Err()
}
