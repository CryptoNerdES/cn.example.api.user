package models

import "net/http"

type HealthzResponse struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

func NewHealthzResponse() *HealthzResponse {
	return &HealthzResponse{
		Message: http.StatusText(http.StatusOK),
		StatusCode: http.StatusOK,
	}
}
