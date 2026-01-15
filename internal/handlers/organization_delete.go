package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

// DeleteOrganization удаляет организацию
func (bh *BaseHandler) DeleteOrganization(w http.ResponseWriter, r *http.Request) {
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

	// Используем репозиторий для удаления организации
	err = bh.orgRepo.DeleteOrganization(id)
	if err != nil {
		http.Error(w, "Failed to delete organization", http.StatusInternalServerError)
		return
	}

	// Set response headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Send response
	if err := json.NewEncoder(w).Encode(map[string]string{"message": "Organization deleted successfully"}); err != nil {
		log.Printf("Error encoding response: %v", err)
	}
}
