package utils

import (
	"net/http"

	"github.com/cperdiansyah/gophermart/internal/api"
	"github.com/go-chi/chi"
)

// ParseURLParamID is a helper to extract an ID from the chi router URL parameters.
// If the ID is empty, it automatically writes a standard 400 Bad Request error response and returns false.
// Otherwise, it returns the string ID and true.
//
// Usage:
//
//	id, ok := utils.ParseURLParamID(w, r, "id")
//	if !ok { return }
func ParseURLParamID(w http.ResponseWriter, r *http.Request, paramKey string) (string, bool) {
	id := chi.URLParam(r, paramKey)
	if id == "" {
		api.ErrorResponse(w, http.StatusBadRequest, "ID produk wajib diisi")
		return "", false
	}
	return id, true
}
