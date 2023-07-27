package utils

import (
	"encoding/json"
	"net/http"
)

// ErrorResponse represents the structure of an error response
type ErrorResponse struct {
	Message string `json:"message"`
}

// SendErrorResponse sends an error response in JSON format
func SendErrorResponse(w http.ResponseWriter, statusCode int, errorMessage string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	errorResponse := ErrorResponse{
		Message: errorMessage,
	}

	json.NewEncoder(w).Encode(errorResponse)
}
