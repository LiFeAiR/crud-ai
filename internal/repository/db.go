package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

// DB структура для работы с базой данных
type DB struct {
	conn *pgxpool.Pool
}

// NewDB создает новый экземпляр подключения к БД
func NewDB(connectionString string) (*DB, error) {
	db, err := pgxpool.New(context.Background(), connectionString)
	if err != nil {
		return nil, fmt.Errorf("failed to create database pool: %w", err)
	}

	// Проверяем соединение
	if err := db.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("Successfully connected to database")
	return &DB{conn: db}, nil
}

// GetConnection возвращает подключение к БД
func (d *DB) GetConnection() *pgxpool.Pool {
	return d.conn
}

// Close закрывает подключение к БД
func (d *DB) Close() error {
	if d.conn != nil {
		d.conn.Close()
	}
	return nil
}
