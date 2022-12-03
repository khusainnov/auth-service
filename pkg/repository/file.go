package repository

import (
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/khusainnov/auth-service/gen/pb"
	"github.com/khusainnov/logging"
	"github.com/lib/pq"
	"google.golang.org/grpc/codes"
)

var (
	logger = logging.GetLogger()
)

const (
	fileTable = "file_table"
)

// TODO: ? ограничение на добавление только одной записи
// TODO: ? update последующих записей

type FilePostgres struct {
	db *sqlx.DB
}

func NewFilePostgres(db *sqlx.DB) *FilePostgres {
	return &FilePostgres{db: db}
}

func (wp *FilePostgres) CreateFile(username string, fileChunks []byte) (*pb.CreateFileResponse, error) {
	var exists bool
	existsQuery := fmt.Sprintf("SELECT exists (SELECT 1 FROM %s WHERE username=$1);", fileTable)
	query := fmt.Sprintf("INSERT INTO %s (username, file_chunks) VALUES ($1, $2);", fileTable)
	updateQuery := fmt.Sprintf("UPDATE %s SET file_chunks=$1 WHERE username=$2;", fileTable)

	if err := wp.db.Get(&exists, existsQuery, username); err != nil {
		return nil, errors.New(err.Error())
	}

	if exists {
		if err := wp.db.QueryRow(updateQuery, fileChunks, username).Err(); err != nil {
			return nil, errors.New(err.Error())
		}

		return &pb.CreateFileResponse{
			Code:    int64(codes.OK),
			Message: "CREATED",
		}, nil
	}

	if err := wp.db.QueryRow(query, username, fileChunks).Err(); err != nil {
		logger.Errorf("%+v", err)
		return nil, errors.New("cannot create file, due to " + err.Error())
	}

	rsp := &pb.CreateFileResponse{
		Code:    int64(codes.OK),
		Message: "CREATED",
	}
	//logger.Infof("%+v", rsp)
	return rsp, nil
}

func (wp *FilePostgres) GetFileByUsername(username string) (*pb.ResponseFile, error) {
	var chunks []byte

	var rspChunks pq.ByteaArray
	query := fmt.Sprintf("SELECT file_chunks FROM %s WHERE username=$1;", fileTable)

	if err := wp.db.Get(&chunks, query, username); err != nil {
		logger.Errorf("%+v", err)
		return nil, err
	}

	rspChunks = append(rspChunks, chunks)
	//fmt.Printf("%+v\n", rspChunks)
	rsp := &pb.ResponseFile{
		Code:     int64(codes.OK),
		Username: username,
		Chunks:   rspChunks,
	}

	//logger.Infof("%+v", rsp)
	return rsp, nil
}

func (wp *FilePostgres) GetAllFiles(username string) (*pb.ResponseFile, error) {
	var admin bool
	fmt.Printf("%s\n", username)
	queryAdmin := fmt.Sprintf("SELECT exists (SELECT 1 FROM %s ut join %s rt on ut.role_id = rt.role_id WHERE (ut.username=$1 and ut.role_id=1) or (ut.email=$1 and ut.role_id=1))", userTable, roleTable)

	if err := wp.db.QueryRow(queryAdmin, username).Scan(&admin); err != nil {
		logger.Errorf("admin: %+v", err)
		return nil, err
	}
	fmt.Printf("%+v\n\n\n", admin)
	if !admin {
		logger.Errorf("%+v", errors.New("you haven't permission for this request"))
		return nil, errors.New("you haven't permission for this request")
	}

	var files pq.ByteaArray

	query := fmt.Sprintf("SELECT file_chunks FROM %s;", fileTable)
	if err := wp.db.Select(&files, query); err != nil {
		logger.Errorf("select: %+v", err)
		return nil, err
	}

	rsp := &pb.ResponseFile{
		Code:     int64(codes.OK),
		Username: username,
		Chunks:   files,
	}

	//logger.Infof("%+v", rsp)

	return rsp, nil
}

type payload struct {
	Name       string `json:"name" db:"name"`
	Surname    string `json:"surname" db:"surname"`
	Patronymic string `json:"patronymic,omitempty" db:"patronymic"`
}

func (wp *FilePostgres) GenerateFile(username string) (*pb.ResponseName, error) {
	var name payload
	query := fmt.Sprintf(`SELECT "name", surname, patronymic FROM %s WHERE username=$1;`, userTable)
	if err := wp.db.Get(&name, query, username); err != nil {
		logger.Errorf("Generate file: " + err.Error())
		return nil, errors.New(err.Error())
	}

	rsp := &pb.ResponseName{
		Name:       name.Name,
		Surname:    name.Surname,
		Patronymic: name.Patronymic,
	}

	return rsp, nil
}
