package auth_service

import (
	"context"
	"errors"
	"fmt"
	"net"
	"sync"

	"github.com/khusainnov/auth-service/gen/pb"
	"github.com/khusainnov/auth-service/pkg/service"
	_ "github.com/khusainnov/logging"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	pb.UnimplementedAuthServiceServer
	sync.Mutex
	srv        *service.Service
	grpcServer *grpc.Server
}

func (s *Server) CreateUser(ctx context.Context, req *pb.User) (*pb.ResponseMsg, error) {
	s.Lock()
	rspMsg, err := s.srv.CreateUser(req)
	s.Unlock()
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return rspMsg, ctx.Err()
}

func (s *Server) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.User, error) {
	rspUser, err := s.srv.GetUser(req)
	if err != nil {
		return &pb.User{}, errors.New(fmt.Sprintf("cannot get user due to %v", err))
	}

	return rspUser, ctx.Err()
}

func (s *Server) UpdateUser(ctx context.Context, req *pb.User) (*pb.ResponseMsg, error) {
	rspMsg, err := s.srv.UpdateUser(req)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("cannot update user due to error: %v", err))
	}

	return rspMsg, ctx.Err()
}

func (s *Server) RunGRPC(port string, srv *service.Service) error {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	s.grpcServer = grpc.NewServer()

	pb.RegisterAuthServiceServer(s.grpcServer, &Server{srv: srv})
	reflection.Register(s.grpcServer)

	return s.grpcServer.Serve(lis)
}
