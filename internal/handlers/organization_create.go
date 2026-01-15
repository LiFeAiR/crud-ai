package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/LiFeAiR/users-crud-ai/internal/models"
)

// CreateOrganization создает новую организацию
func (bh *BaseHandler) CreateOrganization(w http.ResponseWriter, r *http.Request) {
	var org models.Organization

	// Parse JSON request body
	if err := json.NewDecoder(r.Body).Decode(&org); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Используем репозиторий для создания организации
	dbOrg, err := bh.orgRepo.CreateOrganization(&org)
	if err != nil {
		http.Error(w, "Failed to create organization", http.StatusInternalServerError)
		return
	}

	org.ID = dbOrg.ID

	// Set response headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	// Send response
	if err := json.NewEncoder(w).Encode(org); err != nil {
		log.Printf("Error encoding response: %v", err)
	}
}
