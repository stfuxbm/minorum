package utils

import (
	"encoding/json"
	"net/http"

	"github.com/stfuxbm/quote-saints/internal/models"
)

// EncoderJsonData digunakan secara internal untuk encode response JSON ke client.
func EncoderJsonData(w http.ResponseWriter, res models.Response, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(res)
}

// SuccessResponse mengirim response sukses ke client.
func SuccessResponse(w http.ResponseWriter, data interface{}, message string, statusCode int) {
	res := models.Response{
		Success: true,
		Data:    data,
		Message: message,
		Code:    statusCode,
	}
	EncoderJsonData(w, res, statusCode)
}

// ErrorResponse mengirim response error ke client tanpa data.
func ErrorResponse(w http.ResponseWriter, message string, statusCode int) {
	res := models.Response{
		Success: false,
		Message: message,
		Code:    statusCode,
	}
	EncoderJsonData(w, res, statusCode)
}
