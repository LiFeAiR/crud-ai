package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

// DeleteUser общий метод для удаления пользователя
func (bh *BaseHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Получаем userID из query параметров
	userIDStr := r.URL.Query().Get("id")
	if userIDStr == "" {
		http.Error(w, "Missing user ID in query parameters", http.StatusBadRequest)
		return
	}

	// Конвертируем строку в целое число
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Используем репозиторий для удаления пользователя
	if err := bh.userRepo.DeleteUser(userID); err != nil {
		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}

	// Set response headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Send empty response
	if err := json.NewEncoder(w).Encode(map[string]string{"status": "deleted"}); err != nil {
		log.Printf("Error encoding response: %v", err)
	}
}
