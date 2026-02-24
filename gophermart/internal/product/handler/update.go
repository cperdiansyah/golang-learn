package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/cperdiansyah/gophermart/internal/api"
	"github.com/cperdiansyah/gophermart/internal/product/entity"
	"github.com/cperdiansyah/gophermart/internal/utils"
)

func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// 1. Ambil ID dari URL parameter menggunakan logic Utils (SOLID & DRY)
	id, ok := utils.ParseURLParamID(w, r, "id")
	if !ok {
		return
	}

	var req entity.UpdateProductRequest

	// 2. Decode JSON Payload
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		api.ErrorResponse(w, http.StatusBadRequest, "Format JSON tidak valid")
		return
	}

	// 3. Validasi Input Body
	if err := h.validate.Struct(req); err != nil {
		api.ErrorResponse(w, http.StatusBadRequest, "Data input tidak valid")
		return
	}

	// 4. Update data lewat Service
	product, err := h.service.Update(ctx, req, id)
	if err != nil {
		if errors.Is(err, entity.ErrNotFound) {
			api.ErrorResponse(w, http.StatusNotFound, "Produk tidak ditemukan")
			return
		}
		api.ErrorResponse(w, http.StatusInternalServerError, "Gagal mengupdate produk: "+err.Error())
		return
	}

	// 5. Return success
	api.SuccessResponse(w, http.StatusOK, "Berhasil mengupdate produk", product)
}
