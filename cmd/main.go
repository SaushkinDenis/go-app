package main

import (
	todo "github.com/SaushkinDenis/go-app"
	"log"
)

func main() {
	server := new(todo.Server)
	if serverError := server.Run("8000"); serverError != nil {
		log.Fatalf("error occured while running server: %s", serverError.Error())
	}
}
