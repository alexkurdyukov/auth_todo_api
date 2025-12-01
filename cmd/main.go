package main

import (
	"log"
	"todo"
	"todo/pkg/handler"
	"todo/pkg/repository"
	"todo/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	// старт сервера, запускаем и прокидываем в него routes нашего api
	server := new(todo.Server)

	if err := server.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatal("cannot start server")
	}
}
