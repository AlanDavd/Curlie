package main

import (
	"log"

	"curlie/internal/adapter/handler"
)

func main() {
	srv := handler.NewServer()
	if err := srv.Run(":8081"); err != nil {
		log.Fatalf("Failed to start api: %v", err)
	}
}
