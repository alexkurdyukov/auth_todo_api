package main

import (
	"log"
	"todo"
)

func main() {
	server := new(todo.Server)

	if err := server.Run("8080"); err != nil {
		log.Fatal("cannot start server")
	}
}
