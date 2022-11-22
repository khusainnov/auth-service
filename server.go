package auth_service

import (
	"context"
	"net"

	"github.com/khusainnov/auth-service/gen/pb"
	"github.com/khusainnov/auth-service/pkg/endpoint"

	"github.com/khusainnov/logging"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	authService = "authService"
	workService = "workService"
)

var (
	logger = logging.GetLogger()
)

func RunGRPC(port string, as *endpoint.AuthService, ws *endpoint.WorkService) error {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(unaryInterceptor),
		grpc.StreamInterceptor(streamInterceptor),
	)

	pb.RegisterAuthServiceServer(grpcServer, as)
	pb.RegisterWorkServiceServer(grpcServer, ws)
	reflection.Register(grpcServer)

	return grpcServer.Serve(lis)
}

func unaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	logger.Infof("unary interceptor: %+v", info.FullMethod)
	return handler(ctx, req)
}

func streamInterceptor(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	logger.Infof("stream interceptor: %+v", info.FullMethod)
	return handler(srv, stream)
}
