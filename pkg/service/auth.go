package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"os"

	"github.com/khusainnov/auth-service/gen/pb"
	"github.com/khusainnov/auth-service/pkg/repository"
	"github.com/khusainnov/logging"
	"google.golang.org/grpc/codes"
)

type AuthService struct {
	repo repository.Auth
}

var (
	signingKey = os.Getenv("SIGNING_KEY")
	logger     = logging.GetLogger()
)

func NewAuthService(repo repository.Auth) *AuthService {
	return &AuthService{repo: repo}
}

func (as *AuthService) CreateUser(u *pb.User) (*pb.ResponseMsg, error) {
	if u.Username == "" || u.Name == "" || u.Surname == "" || u.Password == "" {
		return &pb.ResponseMsg{Code: int64(codes.InvalidArgument), Message: "EMPTY INPUT DATA"}, errors.New(fmt.Sprintf("Cannot create, due to empty data\n%+v\n", u))
	}
	if !validateEmail(u.GetEmail()) {
		return &pb.ResponseMsg{Code: int64(codes.InvalidArgument), Message: "INCORRECT EMAIL"}, errors.New(fmt.Sprintf("Cannot create, due to incorrect email\n%+v\n", u.Email))
	}

	u.Password = generatePasswordHash(u.Password)

	username, err := as.repo.CreateUser(u)
	if err != nil {
		return &pb.ResponseMsg{Code: int64(codes.AlreadyExists), Message: "Cannot get username"}, err
	}

	if err = sendGreetingMail(username, u.Email); err != nil {
		return nil, errors.New(fmt.Sprintf("cannot send email due to %+v", err))
	}

	return &pb.ResponseMsg{Code: int64(codes.OK), Message: username}, nil
}

// TODO: add reset password

func (as *AuthService) GetUser(login *pb.UserRequest) (*pb.ResponseAuth, error) {
	login.Password = generatePasswordHash(login.Password)
	roleName, err := as.repo.GetUser(login)
	if err != nil {
		return nil, err
	}

	return &pb.ResponseAuth{RoleName: roleName}, nil
}

func (as *AuthService) UpdateUser(u *pb.User) (*pb.ResponseMsg, error) {
	u.Password = generatePasswordHash(u.GetPassword())
	msg, err := as.repo.UpdateUser(u)
	if err != nil {
		return &pb.ResponseMsg{
			Code:    int64(codes.Internal),
			Message: err.Error(),
		}, err
	}

	rspMsg := &pb.ResponseMsg{
		Code:    int64(codes.OK),
		Message: msg,
	}

	return rspMsg, nil
}

func (as *AuthService) DeleteUser(u *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	msg, err := as.repo.DeleteUser(u.GetUsername())
	if err != nil {
		return &pb.DeleteResponse{
			Status:  int64(codes.Internal),
			Message: "cannot delete user due to error | " + msg,
		}, err
	}

	rsp := &pb.DeleteResponse{
		Status:  int64(codes.OK),
		Message: "deleted successfully",
	}

	return rsp, nil
}

func (as *AuthService) ResetPassword(u *pb.UserRequest) (*pb.ResponseMsg, error) {
	pass := u.Password
	u.Password = generatePasswordHash(u.GetPassword())
	fmt.Printf("%+v\n\n", u)
	email, err := as.repo.ResetPassword(u.Username, u.Password)
	if err != nil {
		return &pb.ResponseMsg{
			Code:    int64(codes.Canceled),
			Message: err.Error(),
		}, err
	}
	logger.Infof("%+v", email)

	if err = sendResetPasswordMail(u.Username, email, pass); err != nil {
		return &pb.ResponseMsg{
			Code:    int64(codes.Internal),
			Message: err.Error(),
		}, err
	}

	rsp := &pb.ResponseMsg{
		Code:    int64(codes.OK),
		Message: "UPDATED",
	}

	return rsp, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%+x", hash.Sum([]byte(signingKey)))
}
