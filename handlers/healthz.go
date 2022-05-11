package handlers

import (
	"encoding/json"
	"net/http"

	response "github.com/CryptoNerdES/cn.example.api.user/models/responses"
)

func Healthz(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(response.HealthzResponse{Message: http.StatusText(http.StatusMethodNotAllowed)})
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response.HealthzResponse{Message: http.StatusText(http.StatusOK)})
}
