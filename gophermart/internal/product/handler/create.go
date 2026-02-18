package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cperdiansyah/gophermart/internal/api"
	"github.com/cperdiansyah/gophermart/internal/product/entity"
)

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var req entity.CreateProductRequest

	// 1. Decode JSON Payload
	// Kita decode langsung dari r.Body ke struct req
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		api.ErrorResponse(w, http.StatusBadRequest, "Format Json tidak valid")
		return
	}

	// 2. Validasi Input (menggunakan tag struct yang udah kita buat)
	if err := h.validate.Struct(req); err != nil {
		api.ErrorResponse(w, http.StatusBadRequest, "Format Json tidak valid")
		return
	}

	/* Panggil Service Layer */
	id, err := h.service.CreateProduct(r.Context(), req)
	if err != nil {
		// log error di server
		fmt.Println("error service", err)
		api.ErrorResponse(w, http.StatusInternalServerError, "Gagal membuat product")
		return
	}

	/* Return succes response */
	api.SuccessResponse(w, http.StatusCreated, "Product berhasil dibuat", map[string]int{
		"id": id,
	})
}
