package main

import (
	"log"

	"github.com/alandavd/curlie/internal/infrastructure/server"
)

func main() {
	srv := server.NewServer()
	if err := srv.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
} 