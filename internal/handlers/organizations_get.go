package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

// GetOrganizations получает список организаций с пагинацией
func (bh *BaseHandler) GetOrganizations(w http.ResponseWriter, r *http.Request) {
	// Получаем параметры запроса
	limit := 10
	offset := 0

	if limitStr := r.URL.Query().Get("limit"); limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil {
			limit = l
		}
	}

	if offsetStr := r.URL.Query().Get("offset"); offsetStr != "" {
		if o, err := strconv.Atoi(offsetStr); err == nil {
			offset = o
		}
	}

	// Используем репозиторий для получения организаций
	organizations, err := bh.orgRepo.GetOrganizations(limit, offset)
	if err != nil {
		http.Error(w, "Failed to get organizations", http.StatusInternalServerError)
		return
	}

	// Set response headers
	w.Header().Set("Content-Type", "application/json")

	// Send response
	if err := json.NewEncoder(w).Encode(organizations); err != nil {
		log.Printf("Error encoding response: %v", err)
	}
}
