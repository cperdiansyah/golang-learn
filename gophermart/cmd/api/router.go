package main

import (
	"github.com/cperdiansyah/gophermart/internal/api"
	"github.com/cperdiansyah/gophermart/internal/product/handler"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func NewRouter(pHandler *handler.ProductHandler) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(api.JSONMiddleware)

	// Routes
	// Kita delegate grouping route ke handler masing-masing jika memungkinkan,
	// atau definisikan top-level routing di sini.
	r.Mount("/products", pHandler.Routes()) // Nanti kita update handler biar punya method Routes()

	return r
}
