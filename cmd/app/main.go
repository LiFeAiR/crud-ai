package main

import (
	"os"

	"github.com/LiFeAiR/users-crud-ai/internal/server"
)

func main() {
	// Create and start the server
	s := server.NewServer("8080")

	// Получаем строку подключения к БД из переменной окружения или используем значение по умолчанию
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "host=localhost port=5432 user=postgres password=password dbname=httpserverdb sslmode=disable"
	}

	s.Start(dbURL)
}
