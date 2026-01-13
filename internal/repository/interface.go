package repository

import "github.com/LiFeAiR/users-crud-ai/internal/models"

// UserRepository интерфейс для работы с пользователями
type UserRepository interface {
	CreateUser(user *models.User) (*models.User, error)
	GetUserByID(id int) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id int) error
	GetUsers(limit, offset int) ([]*models.User, error)
	InitDB() error
}
