package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRootHandler(t *testing.T) {
	// Create a test request
	req := httptest.NewRequest("GET", "/", nil)

	// Create a response recorder
	rr := httptest.NewRecorder()

	// Call the handler
	GetRootHandler(rr, req)

	// Check the status code
	assert.Equal(t, http.StatusOK, rr.Code)

	// Check the response body contains expected text
	assert.Contains(t, rr.Body.String(), "Hello, World!")
	assert.Contains(t, rr.Body.String(), "port")
}