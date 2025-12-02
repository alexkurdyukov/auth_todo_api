package main

import (
	"log"
	"todo"
	"todo/pkg/handler"
	"todo/pkg/repository"
	"todo/pkg/service"

	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatal("error while initialize config")
	}

	port := viper.GetString("port")

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	// старт сервера, запускаем и прокидываем в него routes нашего api
	server := new(todo.Server)

	if err := server.Run(port, handlers.InitRoutes()); err != nil {
		log.Fatal("cannot start server")
	}
}

func initConfig() error {
	viper.AddConfigPath("congfigs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
