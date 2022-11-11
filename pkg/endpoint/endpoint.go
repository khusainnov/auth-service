package endpoint

import (
	"context"
	"fmt"

	"github.com/khusainnov/auth-service/gen/pb"
	"github.com/khusainnov/auth-service/pkg/service"
)

type Endpoint struct {
	services *service.Service
}

func NewEndpoint(services *service.Service) *Endpoint {
	return &Endpoint{services: services}
}

func (e *Endpoint) CreateUser(ctx context.Context, req *pb.User) (*pb.SignUpResponse, error) {
	fmt.Printf("Recived: %+v\n", req)
	return &pb.SignUpResponse{Code: 200, Message: fmt.Sprintf("%s\n%s\n", req.GetName(), req.GetSurname())}, nil
}
