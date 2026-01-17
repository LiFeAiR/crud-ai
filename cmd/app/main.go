package main

import (
	"context"
	"log"
	"os"

	"github.com/LiFeAiR/crud-ai/internal/server"
)

func main() {
	log.Fatalf(
		"Failed to serve and listen: %s",
		run().Error(),
	)
}

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Получаем строку подключения к БД из переменной окружения или используем значение по умолчанию
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "host=localhost port=5432 user=postgres password=password dbname=httpserverdb sslmode=disable"
	}

	// Create and start the server
	s := server.NewServer("8080", "2662", dbURL, "your-secret-key")
	return s.Start(ctx)
}
