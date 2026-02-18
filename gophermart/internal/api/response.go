package api

import (
	"encoding/json"
	"net/http"
)

type BaseResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}

/* Helper : JSONContentTypeMiddleware */
// pasang ini di router biar gk perlu nulis w.Header
func JSONMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

/* Helper : ErrorResponse */
// membungkus error menjadi format JSON yang konsisten
func ErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(BaseResponse{
		Message: message,
		Error:   message,
	})
}

/* Helper : SuccessResponse */
// membungkus response sukses menjadi format JSON yang konsisten
func SuccessResponse(w http.ResponseWriter, statusCode int, message string, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(BaseResponse{
		Message: message,
		Data:    data,
	})
}
