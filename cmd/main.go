package main

import (
	"log"
	"os"
	"todo"
	"todo/pkg/handler"
	"todo/pkg/repository"
	"todo/pkg/service"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	if err := initConfig(); err != nil {
		logrus.Fatalf("error while initialize config: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err)
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBname:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		log.Fatalf("could not connect to DB: %s", err)
	}

	defer db.Close()

	port := viper.GetString("port")

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	// старт сервера, запускаем и прокидываем в него routes нашего api
	server := new(todo.Server)

	if err := server.Run(port, handlers.InitRoutes()); err != nil {
		logrus.Fatal("cannot start server")
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
