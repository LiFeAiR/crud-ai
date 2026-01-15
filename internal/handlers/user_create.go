package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/LiFeAiR/users-crud-ai/internal/models"
)

// CreateUser общий метод для создания пользователя
func (bh *BaseHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	// Parse JSON request body
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Используем репозиторий для создания пользователя
	dbUser, err := bh.userRepo.CreateUser(&user)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	user.ID = dbUser.ID

	// Set response headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	// Send response
	if err := json.NewEncoder(w).Encode(user); err != nil {
		log.Printf("Error encoding response: %v", err)
	}
}
