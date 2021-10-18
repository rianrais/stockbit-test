package utilities

import (
	"encoding/json"
	"net/http"
)

// Used to determine error response
type ErrorResponse struct {
	Code    int    `json:"statusCode"`
	Message string `json:"message"`
	Error   string `json:"error"`
}

// This is utility function to return error response to client should the situation called for it
func ErrRes(w http.ResponseWriter, message string, code int, errMsg string) {
	errorResponse := ErrorResponse{
		Code:    code,
		Message: message,
		Error:   errMsg,
	}
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(errorResponse)
	return
}
