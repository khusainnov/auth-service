package main

import (
	"context"
	"os"

	"github.com/joho/godotenv"
	as "github.com/khusainnov/auth-service"
	"github.com/khusainnov/auth-service/driver"
	"github.com/khusainnov/auth-service/pkg/endpoint"
	"github.com/khusainnov/auth-service/pkg/repository"
	"github.com/khusainnov/auth-service/pkg/service"
	"github.com/khusainnov/logging"
)

var (
	logger = logging.GetLogger()
	ctx    = context.Background()
)

func main() {
	logger.Infoln("Initialization config")
	if err := godotenv.Load("./config/.env"); err != nil {
		logger.Fatalf("Error due load config: %s", err.Error())
	}

	// Postgres DB Connect
	logger.Infoln("connecting to postgres")
	db, err := driver.NewPostgresDB(driver.ConfigPG{
		Host:     os.Getenv("PG_HOST"),
		Port:     os.Getenv("PG_PORT"),
		User:     os.Getenv("PG_USER"),
		DBName:   os.Getenv("PG_DB_NAME"),
		SSLMode:  os.Getenv("PG_SSL_MODE"),
		Password: os.Getenv("PG_PASSWORD"),
	})
	if err != nil {
		logger.Errorf("cannot run postgres db, due to error: %s", err.Error())
	}

	//Initializing layers
	repo := repository.NewRepository(db)
	services := service.NewService(repo)
	authService := endpoint.NewAuthService(services)
	fileService := endpoint.NewWorkService(services)
	statisticsService := endpoint.NewStatisticsService(services)

	//starting grpc server
	logger.Infof("Starting server on port: %s", os.Getenv("PORT"))
	if err = as.RunGRPC(os.Getenv("PORT"), authService, fileService, statisticsService); err != nil {
		logger.Errorf("Error due start the server: %s", err.Error())
	}
}
