package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/CryptoNerdES/cn.example.api.user/handlers"
	"github.com/CryptoNerdES/cn.example.api.user/models"
	"github.com/stretchr/testify/assert"
)

func TestReturnStatusOKWhenHealthz(t *testing.T) {
	// Test set up
	var response models.HealthzResponse
	expectedResponse := models.NewHealthzResponse()
	req := httptest.NewRequest("GET", "/healthz", nil)
	w := httptest.NewRecorder()

	// Execute the function to check
	handlers.Healthz(w, req)
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Error(err)
	}

	// Check if function is doing the work right
	assert.Equal(t, http.StatusOK, w.Result().StatusCode)
	assert.Equal(t, expectedResponse.Message, response.Message)
	assert.Equal(t, expectedResponse.StatusCode, response.StatusCode)
}
