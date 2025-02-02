package helpers

import (
	"encoding/json"
	"net/http"
)

// APIResponse is the structure used for standardized API responses
type APIResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Code    int         `json:"code"`
}

// Respond writes the standardized response to the HTTP response writer
func Respond(w http.ResponseWriter, status bool, message string, data interface{}, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	response := APIResponse{
		Status:  status,
		Message: message,
		Data:    data,
		Code:    code,
	}
	json.NewEncoder(w).Encode(response)
}
