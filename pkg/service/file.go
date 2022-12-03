package service

import (
	"github.com/khusainnov/auth-service/gen/pb"
	"github.com/khusainnov/auth-service/pkg/repository"
)

type FileService struct {
	repo repository.File
}

func NewFileService(repo repository.File) *FileService {
	return &FileService{repo: repo}
}

func (ws *FileService) CreateFile(file *pb.UserFile) (*pb.CreateFileResponse, error) {
	rsp, err := ws.repo.CreateFile(file.Username, file.File)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

func (ws *FileService) GetFileByUsername(req *pb.FileRequest) (*pb.ResponseFile, error) {
	rsp, err := ws.repo.GetFileByUsername(req.Username)
	if err != nil {
		return nil, err
	}

	return rsp, nil
}

func (ws *FileService) GetAllFiles(req *pb.FileRequest) (*pb.ResponseFile, error) {
	rsp, err := ws.repo.GetAllFiles(req.Username)
	if err != nil {
		return nil, err
	}

	return rsp, nil
}

func (ws *FileService) GenerateFile(req *pb.FileRequest) (*pb.ResponseName, error) {
	rsp, err := ws.repo.GenerateFile(req.Username)
	if err != nil {
		return nil, err
	}

	return rsp, nil
}

/*func CreateFile(file *pb.UserWork) (*pb.ResponseWork, error) {
	var bb bytes.Buffer
	bb.Write(file.GetFile())

	ud, _ := uuid.NewRandom()

	f, _ := os.Create("pdf/" + ud.String() + ".jpeg")

	_, err := bb.WriteTo(f)

	if err != nil {
		return &pb.ResponseWork{Code: int64(codes.Internal)}, errors.New("cannot write bytes into pdf file")
	}
	cmd := exec.Command("sudo", "convert", fmt.Sprintf("/tmp/hackaton/pdf/%s.jpeg", ud.String()), fmt.Sprintf("/tmp/hackaton/pdf/%s.pdf", ud.String()))
	if err = cmd.Start(); err != nil {
		return &pb.ResponseWork{Code: int64(codes.Internal)}, errors.New(fmt.Sprintf("cannot convert from jpeg to pdf file\n%s\n", err.Error()))
	}
	rspBytes, err := ioutil.ReadFile("pdf/" + ud.String() + ".jpeg")
	if err != nil {
		return &pb.ResponseWork{Code: int64(codes.Internal)}, errors.New(fmt.Sprintf("cannot read bytes from pdf file\n%s\n", err.Error()))
	}

	return &pb.ResponseWork{Code: int64(codes.OK), File: rspBytes}, nil
}*/
