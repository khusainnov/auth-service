package service

import (
	"github.com/khusainnov/auth-service/gen/pb"
	"github.com/khusainnov/auth-service/pkg/repository"
)

type StatisticsService struct {
	repo repository.Statistics
}

func NewStatisticsService(repo repository.Statistics) *StatisticsService {
	return &StatisticsService{repo: repo}
}

func (s *StatisticsService) GetStatistics(req *pb.StatisticsRequest) (*pb.StatisticsResponse, error) {
	rsp, err := s.repo.GetStatistics(req.Username)
	if err != nil {
		return rsp, err
	}

	return rsp, nil
}
