package main

import (
	"log"
	"net/http"

	"github.com/CryptoNerdES/cn.example.api.user/handlers"
)

func main() {
	http.HandleFunc("/healthz", handlers.Healthz)

	log.Fatal(http.ListenAndServe(":10000", nil))
}