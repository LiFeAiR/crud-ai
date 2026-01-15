package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/LiFeAiR/users-crud-ai/internal/models"
)

// UpdateUser общий метод для обновления пользователя
func (bh *BaseHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	// Parse JSON request body
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Используем репозиторий для обновления пользователя
	if err := bh.userRepo.UpdateUser(&user); err != nil {
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	// Set response headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Send response
	if err := json.NewEncoder(w).Encode(user); err != nil {
		log.Printf("Error encoding response: %v", err)
	}
}
