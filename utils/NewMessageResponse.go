package utils

import (
	"encoding/json"
	"net/http"
)

type ResponseMessage struct {
	Message string `json:"message"`
}

func NewMessageResponse(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := ResponseMessage{
		Message: message,
	}

	json.NewEncoder(w).Encode(response)
}
