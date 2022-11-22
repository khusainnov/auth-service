package main

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	as "github.com/khusainnov/auth-service"
	"github.com/khusainnov/auth-service/driver"
	"github.com/khusainnov/auth-service/pkg/endpoint"
	"github.com/khusainnov/auth-service/pkg/repository"
	"github.com/khusainnov/auth-service/pkg/service"
	"github.com/khusainnov/logging"
)

// TODO: регистрация пользователя
// TODO: сделать docker-file и залить на сервер
// TODO: написать Ильназу чтобы добавил клиента для отправки данных в бота
// TODO: проверка регистрации через бота

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

	// Redis DB Connect
	logger.Infoln("connecting to redis")
	rdb, err := driver.NewRedisDB(
		driver.ConfigRedis{
			// if you start with docker-compose, change from localhost to Getenv
			Port:     fmt.Sprintf("%s:%s" /*os.Getenv("REDIS_NAME")*/, "localhost", os.Getenv("REDIS_PORT")),
			Password: os.Getenv("REDIS_PASSWORD"),
			DB:       0,
		},
		ctx,
	)
	if err != nil {
		logger.Fatalf("cannot connect to redis, due to error: %s", err.Error())
	}

	//Initializing layers
	repo := repository.NewRepository(db, rdb)
	services := service.NewService(repo)
	authService := endpoint.NewAuthService(services)
	workService := endpoint.NewWorkService(services)

	//starting grpc server
	logger.Infof("Starting server on port: %s", os.Getenv("PORT"))
	if err = as.RunGRPC(os.Getenv("PORT"), authService, workService); err != nil {
		logger.Errorf("Error due start the server: %s", err.Error())
	}
}
