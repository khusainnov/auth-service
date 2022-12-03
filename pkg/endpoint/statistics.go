package endpoint

import (
	"context"

	"github.com/khusainnov/auth-service/gen/pb"
	"github.com/khusainnov/auth-service/pkg/service"
)

type StatisticsService struct {
	pb.UnimplementedStatServiceServer
	services *service.Service
}

func NewStatisticsService(services *service.Service) *StatisticsService {
	return &StatisticsService{services: services}
}

func (s *StatisticsService) GetStatistics(ctx context.Context, req *pb.StatisticsRequest) (*pb.StatisticsResponse, error) {
	logger.Printf("%+v\n", req)
	rsp, err := s.services.GetStatistics(req)
	if err != nil {
		logger.Printf("%+v\n", err)
		return rsp, err
	}

	return rsp, ctx.Err()
}
