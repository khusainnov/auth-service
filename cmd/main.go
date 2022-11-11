package main

import (
	"os"

	"github.com/joho/godotenv"
	as "github.com/khusainnov/auth-service"
	"github.com/khusainnov/auth-service/driver"
	"github.com/khusainnov/auth-service/pkg/repository"
	"github.com/khusainnov/auth-service/pkg/service"
	"github.com/khusainnov/logging"
	"github.com/sirupsen/logrus"
)

// TODO: регистрация пользователя
// TODO: сделать docker-file и залить на сервер
// TODO: написать Ильназу чтобы добавил клиента для отправки данных в бота
// TODO: проверка регистрации через бота

var (
	logger = logging.GetLogger()
)

func main() {
	logger.Infoln("Initialization config")
	if err := godotenv.Load("./config/.env"); err != nil {
		logger.Fatalf("Error due load config: %s", err.Error())
	}

	// DB Connect
	logger.Infoln("Initializing PostgresDB")
	db, err := driver.NewPostgresDB(driver.ConfigPG{
		Host:     os.Getenv("PG_HOST"),
		Port:     os.Getenv("PG_PORT"),
		User:     os.Getenv("PG_USER"),
		DBName:   os.Getenv("PG_DB_NAME"),
		SSLMode:  os.Getenv("PG_SSL_MODE"),
		Password: os.Getenv("PG_PASSWORD"),
	})
	if err != nil {
		logrus.Errorf("Cannot run db, due to error: %s", err.Error())
	}

	repo := repository.NewRepository(db)
	services := service.NewService(repo)
	//endpoints := enpoint.NewEnpoint(services)

	// Redis connect

	//starting grpc server
	logger.Infof("Starting server on port: %s", os.Getenv("PORT"))
	s := as.Server{}
	if err = s.RunGRPC(os.Getenv("PORT"), services); err != nil {
		logger.Errorf("Error due start the server: %s", err.Error())
	}
}
