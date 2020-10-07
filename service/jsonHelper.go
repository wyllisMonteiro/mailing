package service

import (
	"encoding/json"
	"net/http"
)

// ErrorResponse : Structure format of JSON return while error occured
type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// WriteJSON : Return JSON if there is no error
func WriteJSON(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	data, _ := json.Marshal(v)
	w.Write(data)
}

// WriteErrorJSON : Return JSON if there is an error
func WriteErrorJSON(w http.ResponseWriter, status int, message string) {
	WriteJSON(w, status, ErrorResponse{
		Status:  "error",
		Message: message,
	})
}
