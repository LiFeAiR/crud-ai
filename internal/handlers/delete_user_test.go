package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteUserHandler(t *testing.T) {
	// Create a test request
	req := httptest.NewRequest("DELETE", "/user?id=1", nil)

	// Создаем мок репозиторий
	mockRepo := new(MockUserRepository)

	// Определяем ожидаемое поведение мока
	mockRepo.On("DeleteUser", 1).Return(nil)

	// Создаем базовый обработчик с моком
	baseHandler := &BaseHandler{
		userRepo: mockRepo,
	}

	// Create a response recorder
	rr := httptest.NewRecorder()

	// Call the handler
	baseHandler.DeleteUser(rr, req)

	// Check the status code
	assert.Equal(t, http.StatusOK, rr.Code)

	// Check the response body
	var response map[string]string
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "deleted", response["status"])
}

func TestDeleteUserHandler_InvalidID(t *testing.T) {
	// Create a test request with invalid ID
	req := httptest.NewRequest("DELETE", "/user?id=abc", nil)

	// Создаем мок репозиторий
	mockRepo := new(MockUserRepository)

	// Создаем базовый обработчик с моком
	baseHandler := &BaseHandler{
		userRepo: mockRepo,
	}

	// Create a response recorder
	rr := httptest.NewRecorder()

	// Call the handler
	baseHandler.DeleteUser(rr, req)

	// Check the status code
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestDeleteUserHandler_UserNotFound(t *testing.T) {
	// Create a test request
	req := httptest.NewRequest("DELETE", "/user?id=1", nil)

	// Создаем мок репозиторий
	mockRepo := new(MockUserRepository)

	// Определяем ожидаемое поведение мока - возвращаем ошибку
	mockRepo.On("DeleteUser", 1).Return(errors.New("user not found"))

	// Создаем базовый обработчик с моком
	baseHandler := &BaseHandler{
		userRepo: mockRepo,
	}

	// Create a response recorder
	rr := httptest.NewRecorder()

	// Call the handler
	baseHandler.DeleteUser(rr, req)

	// Check the status code
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}
