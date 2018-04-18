package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func respond(w http.ResponseWriter, value interface{}, statusCode int) {
	w.Header().Add(headerContentType, contentTypeJSON)
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(value); err != nil {
		log.Printf("error encoding JSON: %v", err)
	}
}
