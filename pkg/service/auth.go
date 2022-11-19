package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"os"

	"github.com/khusainnov/auth-service/gen/pb"
	"github.com/khusainnov/auth-service/pkg/repository"
	"google.golang.org/grpc/codes"
)

type AuthService struct {
	repo repository.Auth
}

var (
	signingKey = os.Getenv("SIGNING_KEY")
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

	token, err := GenerateToken(username)
	if err != nil {
		return &pb.ResponseMsg{Code: int64(codes.Internal), Message: "Cannot get username"}, errors.New(fmt.Sprintf("error due creating token: %s\n", err.Error()))
	}

	if err = sendMail(token, u.Email); err != nil {
		return nil, errors.New(fmt.Sprintf("cannot send email due to %+v", err))
	}

	return &pb.ResponseMsg{Code: int64(codes.OK), Message: token}, nil
}

// TODO: add role table
// TODO: add permissions by role

// TODO: get user and send token to client
// TODO: update user, before update check token
// TODO: add reset password

func (as *AuthService) GetUser(login *pb.UserRequest) (*pb.ResponseToken, error) {
	login.Password = generatePasswordHash(login.Password)
	user, err := as.repo.GetUser(login)
	if err != nil {
		return nil, err
	}

	token, err := GenerateToken(user.Username)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error due creating token: %s\n", err.Error()))
	}

	return &pb.ResponseToken{Token: token}, nil
}
func (as *AuthService) UpdateUser(u *pb.User) (*pb.ResponseMsg, error) {
	u.Password = generatePasswordHash(u.GetPassword())
	return nil, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%+x", hash.Sum([]byte(signingKey)))
}
