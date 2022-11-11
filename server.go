package auth_service

import (
	"context"
	"errors"
	"fmt"
	"net"

	"github.com/khusainnov/auth-service/gen/pb"
	"github.com/khusainnov/auth-service/pkg/service"
	_ "github.com/khusainnov/logging"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	pb.UnimplementedAuthServiceServer
	srv        *service.Service
	grpcServer *grpc.Server
}

func (s *Server) CreateUser(ctx context.Context, req *pb.User) (*pb.SignUpResponse, error) {
	fmt.Printf("Recived: %+v\n", req)
	id, err := s.srv.CreateUser(req)
	if err != nil {
		return nil, errors.New("cannot create user")
	}
	return &pb.SignUpResponse{Code: 200, Message: fmt.Sprintf("Created with id: %d\n", id)}, nil
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
