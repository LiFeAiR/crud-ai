package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/LiFeAiR/users-crud-ai/internal/models"
)

// UpdateOrganization обновляет информацию об организации
func (bh *BaseHandler) UpdateOrganization(w http.ResponseWriter, r *http.Request) {
	// Получаем ID из query параметров
	iDStr := r.URL.Query().Get("id")
	if iDStr == "" {
		http.Error(w, "Missing organization ID in query parameters", http.StatusBadRequest)
		return
	}

	// Конвертируем строку в целое число
	id, err := strconv.Atoi(iDStr)
	if err != nil {
		http.Error(w, "Invalid organization ID", http.StatusBadRequest)
		return
	}

	var org models.Organization

	// Parse JSON request body
	if err := json.NewDecoder(r.Body).Decode(&org); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Устанавливаем ID из URL
	org.ID = id

	// Используем репозиторий для обновления организации
	err = bh.orgRepo.UpdateOrganization(&org)
	if err != nil {
		http.Error(w, "Failed to update organization", http.StatusInternalServerError)
		return
	}

	// Set response headers
	w.Header().Set("Content-Type", "application/json")

	// Send response
	if err := json.NewEncoder(w).Encode(org); err != nil {
		log.Printf("Error encoding response: %v", err)
	}
}
