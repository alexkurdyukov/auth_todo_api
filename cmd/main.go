package main

import (
	"log"
	"todo"
	"todo/pkg/handler"
)

func main() {
	// инициализируем handlers и роуты
	handler := new(handler.Handler)
	router := handler.InitRoutes()

	// старт сервера, запускаем и прокидываем в него routes нашего api
	server := new(todo.Server)

	if err := server.Run("8080", router); err != nil {
		log.Fatal("cannot start server")
	}
}
