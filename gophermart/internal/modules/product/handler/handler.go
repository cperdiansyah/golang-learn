package handler

import (
	"github.com/cperdiansyah/gophermart/internal/modules/product/service"
	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"
)

// import "github.com/cperdiansyah/gophermart/internal/modules/product/service"

// product handler struct menyimpan deps yang dibutuhkan
type ProductHandler struct {
	service  service.Service
	validate *validator.Validate
}

/* new product handler bikin instance handler baru dan setup routing */
func NewProductHandler(s service.Service, v *validator.Validate) *ProductHandler {
	return &ProductHandler{
		service:  s,
		validate: v,
	}
}

func (h *ProductHandler) Routes() chi.Router {
	r := chi.NewRouter()
	r.Post("/", h.CreateProduct)
	r.Get("/", h.GetAllProduct)
	r.Get("/{id}", h.GetProductByID)
	r.Put("/{id}", h.UpdateProduct)
	r.Delete("/{id}", h.DeleteProduct)
	return r
}
