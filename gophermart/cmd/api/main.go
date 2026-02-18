package main

import (
	"log"
	"net/http"

	"github.com/cperdiansyah/gophermart/internal/infrastructure"
	"github.com/cperdiansyah/gophermart/internal/product/handler"
	"github.com/cperdiansyah/gophermart/internal/product/repository"
	"github.com/cperdiansyah/gophermart/internal/product/service"
	"github.com/go-playground/validator/v10"
)

func main() {
	// 1. Setup Configuration & Infrastructure
	// (Idealnya connStr dari env var/config file)
	connStr := "postgres://postgres:password@localhost:5432/gophermart?sslmode=disable"
	db, err := infrastructure.NewPostgresDB(connStr)
	if err != nil {
		log.Fatalf("Infrastructure failed: %v", err)
	}
	defer db.Close()
	log.Println("✅ Database connected")

	// 2. Wiring Dependencies (Dependency Injection)
	// Layer Data
	pRepo := repository.NewPostgresRepository(db)

	// Layer Business
	pService := service.NewProductService(pRepo)

	// Layer Validator (3rd party)
	v := validator.New()

	// Layer Presentation (Handler)
	pHandler := handler.NewProductHandler(pService, v) // Perhatikan: NewProductHandler perlu disesuaikan return struct-nya saja, hapus setup router di constructor lama jika ada

	// 3. Setup Router
	r := NewRouter(pHandler) // Panggil fungsi dari router.go (package main)

	// 4. Start Server
	port := ":8080"
	log.Printf("🚀 Server running on http://localhost%s", port)
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatalf("Server crash: %v", err)
	}
}
