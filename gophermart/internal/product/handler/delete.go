package handler

import (
	"errors"
	"net/http"

	"github.com/cperdiansyah/gophermart/internal/api"
	"github.com/cperdiansyah/gophermart/internal/product/entity"
	"github.com/cperdiansyah/gophermart/internal/utils"
)

func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// 1. Ambil ID dari URL parameter menggunakan custom utils (DRY principle)
	id, ok := utils.ParseURLParamID(w, r, "id")
	if !ok {
		return
	}

	// 2. Delete data lewat Service
	deletedID, err := h.service.Delete(ctx, id)
	if err != nil {
		if errors.Is(err, entity.ErrNotFound) {
			api.ErrorResponse(w, http.StatusNotFound, "Produk tidak ditemukan")
			return
		}
		api.ErrorResponse(w, http.StatusInternalServerError, "Gagal menghapus produk: "+err.Error())
		return
	}

	// 3. Return success
	api.SuccessResponse(w, http.StatusOK, "Berhasil menghapus produk", map[string]string{
		"id": deletedID,
	})
}
