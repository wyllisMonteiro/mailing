package service

import (
	"net/http"
	"encoding/json"
)

type ErrorResponse struct {
	Status string `json:"status"`
	Message string `json:"message"`
}

func WriteJSON(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	data, _ := json.Marshal(v)
	w.Write(data)
}

func WriteErrorJSON(w http.ResponseWriter, status int, message string) {
	WriteJSON(w, status, ErrorResponse{
		Status: "error",
		Message: message,
	})
}