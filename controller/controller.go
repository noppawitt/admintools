package controller

import (
	"encoding/json"
	"net/http"
)

// Controller is a base controller
type Controller struct{}

// JSON renders JSON response
func (c *Controller) JSON(w http.ResponseWriter, code int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(data)
	return err
}

// Error render json error message
func (c *Controller) Error(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	type ErrorMessage struct {
		Message string `json:"message"`
	}
	errorMessage := ErrorMessage{message}
	json.NewEncoder(w).Encode(errorMessage)
}
