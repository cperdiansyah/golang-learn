package handler

import (
	"net/http"
	"strconv"

	"github.com/cperdiansyah/gophermart/internal/api"
	"github.com/cperdiansyah/gophermart/internal/modules/product/entity"
	"github.com/cperdiansyah/gophermart/internal/utils"
)

type GetAllQuery struct {
	Limit int `validate:"gte=1,lte=100"`
	Page  int `validate:"gte=1"`
}

func (h *ProductHandler) GetAllProduct(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// 1. Setup Default Pagination
	query := GetAllQuery{
		Limit: 10,
		Page:  1,
	}

	// 2. Parse Query Parameters (Robust Validation)
	if limitStr := r.URL.Query().Get("limit"); limitStr != "" {
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			api.ErrorResponse(w, http.StatusBadRequest, "Parameter 'limit' harus berupa angka")
			return
		}
		query.Limit = limit
	}

	if pageStr := r.URL.Query().Get("page"); pageStr != "" {
		page, err := strconv.Atoi(pageStr)
		if err != nil {
			api.ErrorResponse(w, http.StatusBadRequest, "Parameter 'page' harus berupa angka")
			return
		}
		query.Page = page
	}

	// 3. Validasi batasan Limit & Page menggunakan validator
	if err := h.validate.Struct(query); err != nil {
		api.ErrorResponse(w, http.StatusBadRequest, "Parameter pagination tidak valid (limit: 1-100, page: >= 1)")
		return
	}

	// 4. Hitung Offset
	offset := (query.Page - 1) * query.Limit

	// 5. Ambil data dari Service dengan Limit & Offset
	products, err := h.service.FindAll(ctx, query.Limit, offset)
	if err != nil {
		api.ErrorResponse(w, http.StatusInternalServerError, "Gagal mengambil data produk")
		return
	}

	// Handle data kosong
	if len(products) == 0 {
		api.SuccessResponse(w, http.StatusOK, "Data produk kosong", []entity.Product{})
		return
	}

	// 6. Return response sukses
	api.SuccessResponse(w, http.StatusOK, "Berhasil mengambil data produk", products)
}

func (h *ProductHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// 1. Ambil ID dari URL parameter menggunakan custom utils (DRY principle)
	id, ok := utils.ParseURLParamID(w, r, "id")
	if !ok {
		return
	}

	// 2. Ambil data dari Service
	product, err := h.service.FindByID(ctx, id)
	if err != nil {
		api.ErrorResponse(w, http.StatusInternalServerError, "Gagal mengambil data produk")
		return
	}

	// 3. Handle data tidak ditemukan (repository return nil, nil)
	if product == nil {
		api.ErrorResponse(w, http.StatusNotFound, "Produk tidak ditemukan")
		return
	}

	// 4. Return response sukses
	api.SuccessResponse(w, http.StatusOK, "Berhasil mengambil data produk", product)
}
