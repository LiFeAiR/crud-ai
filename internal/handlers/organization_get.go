package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

// GetOrganization получает организацию по ID
func (bh *BaseHandler) GetOrganization(w http.ResponseWriter, r *http.Request) {
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

	// Используем репозиторий для получения организации
	org, err := bh.orgRepo.GetOrganizationByID(id)
	if err != nil {
		http.Error(w, "Organization not found", http.StatusNotFound)
		return
	}

	// Set response headers
	w.Header().Set("Content-Type", "application/json")

	// Send response
	if err := json.NewEncoder(w).Encode(org); err != nil {
		log.Printf("Error encoding response: %v", err)
	}
}
