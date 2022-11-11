package main

import (
	"os"

	"github.com/joho/godotenv"
	as "github.com/khusainnov/auth-service"
	"github.com/khusainnov/logging"
)

var (
	logger = logging.GetLogger()
)

func main() {
	logger.Infoln("Initialization config")
	if err := godotenv.Load("./config/.env"); err != nil {
		logger.Fatalf("Error due load config: %s", err.Error())
	}

	// DB Connect

	// Redis connect

	//starting grpc server
	logger.Infof("Starting server on port: %s", os.Getenv("PORT"))
	s := as.Server{}
	if err := s.RunGRPC(os.Getenv("PORT")); err != nil {
		logger.Errorf("Error due start the server: %s", err.Error())
	}
}
