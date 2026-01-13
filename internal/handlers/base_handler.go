package handlers

import (
	"github.com/LiFeAiR/users-crud-ai/internal/repository"
)

// BaseHandler базовый обработчик, который принимает репозиторий
type BaseHandler struct {
	userRepo repository.UserRepository
}

// NewBaseHandler создает новый базовый обработчик
func NewBaseHandler(userRepo repository.UserRepository) *BaseHandler {
	return &BaseHandler{
		userRepo: userRepo,
	}
}
