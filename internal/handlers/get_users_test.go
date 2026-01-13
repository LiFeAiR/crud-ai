package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUsersHandler(t *testing.T) {
	// Создаем тестовый сервер
	server := setupTestServer(t)
	defer server.Close()

	// Тест 1: Получение списка пользователей
	t.Run("GetUsers", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/users?limit=10&offset=0", nil)
		w := httptest.NewRecorder()
		server.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})
}