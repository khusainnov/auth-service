package endpoint

import (
	"context"

	"github.com/khusainnov/auth-service/gen/pb"
	"github.com/khusainnov/auth-service/pkg/service"
)

type WorkService struct {
	pb.UnimplementedWorkServiceServer
	services *service.Service
}

func NewWorkService(services *service.Service) *WorkService {
	return &WorkService{services: services}
}

func (ws *WorkService) CreateFile(ctx context.Context, req *pb.UserWork) (*pb.ResponseWork, error) {
	rsp, err := ws.services.CreateFile(req)
	if err != nil {
		return nil, err
	}

	return rsp, ctx.Err()
}
