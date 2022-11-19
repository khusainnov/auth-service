package endpoint

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/khusainnov/auth-service/gen/pb"
	"github.com/khusainnov/auth-service/pkg/service"
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
	s.Lock()
	rspMsg, err := s.srv.CreateUser(req)
	s.Unlock()
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return rspMsg, ctx.Err()
}

func (s *AuthService) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.ResponseToken, error) {
	rspUser, err := s.srv.GetUser(req)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("cannot get user due to %v", err))
	}

	return rspUser, ctx.Err()
}

func (s *AuthService) UpdateUser(ctx context.Context, req *pb.User) (*pb.ResponseMsg, error) {
	rspMsg, err := s.srv.UpdateUser(req)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("cannot update user due to error: %v", err))
	}

	return rspMsg, ctx.Err()
}

func (s *AuthService) ResetPassword(ctx context.Context, req *pb.UserRequest) (*pb.ResponseMsg, error) {
	return nil, nil
}
