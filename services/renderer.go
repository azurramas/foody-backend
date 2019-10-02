package services

import (
	"encoding/json"
	"net/http"
)

//WriteJSON encodes interface to json and writes it to writer
func WriteJSON(w http.ResponseWriter, data interface{}, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}
