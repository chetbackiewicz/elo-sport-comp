package services

import (
	"encoding/json"
	"net/http"
)

// ErrorResponse represents a standardized error response structure
type ErrorResponse struct {
	Message string `json:"message"`
}

// AuthResponse represents an authentication response
type AuthResponse struct {
	Message string `json:"message"`
	IsAuth  bool   `json:"isAuth"`
}

// SendError sends a JSON error response with the specified message and status code
func SendError(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(ErrorResponse{Message: message})
}

// SendJSON sends a JSON response with the specified data
func SendJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
