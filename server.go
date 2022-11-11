package auth_service

import (
	"context"
	"fmt"
	"net"

	"github.com/khusainnov/auth-service/gen/pb"
	"github.com/khusainnov/logging"
	"google.golang.org/grpc"
)

var (
	logger = logging.GetLogger()
)

type Server struct {
	pb.UnimplementedAuthServiceServer
	grpcServer *grpc.Server
}

func (s *Server) CreateUser(ctx context.Context, req *pb.User) (*pb.SignUpResponse, error) {
	fmt.Printf("Recived: %+v\n", req)
	return &pb.SignUpResponse{Code: 200, Message: fmt.Sprintf("%s\n%s\n", req.GetName(), req.GetSurname())}, nil
}

func (s *Server) RunGRPC(port string) error {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	s.grpcServer = grpc.NewServer()

	pb.RegisterAuthServiceServer(s.grpcServer, &Server{})

	return s.grpcServer.Serve(lis)
}
