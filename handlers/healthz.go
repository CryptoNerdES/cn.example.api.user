package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/CryptoNerdES/cn.example.api.user/models"
)

func Healthz(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.NewHealthzResponse())
}