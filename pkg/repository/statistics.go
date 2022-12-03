package repository

import (
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/khusainnov/auth-service/gen/pb"
	"google.golang.org/grpc/codes"
)

type StatisticsPostgres struct {
	db *sqlx.DB
}

func NewStatisticsPostgres(db *sqlx.DB) *StatisticsPostgres {
	return &StatisticsPostgres{db: db}
}

type payloadS struct {
	UserNumb int `json:"user_numb" db:"user_numb"`
	FileNumb int `json:"file_numb" db:"file_numb"`
}

func (s *StatisticsPostgres) GetStatistics(username string) (*pb.StatisticsResponse, error) {
	var admin bool
	var resp payloadS
	query := fmt.Sprintf("SELECT user_numb, file_numb FROM %s;", statisticsTable)
	queryAdmin := fmt.Sprintf("SELECT exists (SELECT 1 FROM %s ut join %s rt on ut.role_id = rt.role_id WHERE (ut.username=$1 and ut.role_id=1) or (ut.email=$1 and ut.role_id=1))", userTable, roleTable)

	if err := s.db.QueryRow(queryAdmin, username).Scan(&admin); err != nil {
		logger.Errorf("admin: %+v", err)
		return nil, err
	}

	fmt.Printf("%+v\n\n\n", admin)
	if !admin {
		logger.Errorf("%+v", errors.New("you haven't permission for this request"))
		return nil, errors.New("you haven't permission for this request")
	}

	if err := s.db.Get(&resp, query); err != nil {
		return &pb.StatisticsResponse{
			Code:     int64(codes.Internal),
			Message:  err.Error(),
			UserNumb: 0,
			FileNumb: 0,
		}, err
	}
	fmt.Printf("STATISTICS: %+v\n", resp)

	return &pb.StatisticsResponse{
		Code:     int64(codes.OK),
		Message:  "successfully",
		UserNumb: int64(resp.UserNumb),
		FileNumb: int64(resp.FileNumb),
	}, nil
}
