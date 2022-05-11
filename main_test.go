package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/CryptoNerdES/cn.example.api.user/handlers"
	response "github.com/CryptoNerdES/cn.example.api.user/models/responses"
	"github.com/stretchr/testify/assert"
)

func TestReturnStatusOKWhenHealthz(t *testing.T) {
	// Test set up
	var res response.HealthzResponse
	expectedResponse := response.NewHealthzResponse()
	req := httptest.NewRequest("GET", "/healthz", nil)
	w := httptest.NewRecorder()

	// Execute the function to check
	handlers.Healthz(w, req)
	err := json.Unmarshal(w.Body.Bytes(), &res)
	if err != nil {
		t.Error(err)
	}

	// Check if function is doing the work right
	assert.Equal(t, http.StatusOK, w.Result().StatusCode)
	assert.Equal(t, expectedResponse.Message, res.Message)
	assert.Equal(t, expectedResponse.StatusCode, res.StatusCode)
}
