package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/cperdiansyah/gophermart/internal/config"
	"github.com/cperdiansyah/gophermart/internal/infrastructure"
	"github.com/cperdiansyah/gophermart/internal/modules/product/handler"
	"github.com/cperdiansyah/gophermart/internal/modules/product/repository"
	"github.com/cperdiansyah/gophermart/internal/modules/product/service"
	"github.com/go-playground/validator/v10"
)

func main() {
	// 1. Load Configurations
	cfg := config.LoadConfig()

	// 2. Setup Infrastructure
	db, err := infrastructure.NewPostgresDB(cfg.DBConnStr)
	if err != nil {
		log.Fatalf("Infrastructure failed: %v", err)
	}
	defer db.Close()
	log.Println("✅ Database connected")

	// 3. Wiring Dependencies (Dependency Injection)
	// Layer Data
	pRepo := repository.NewPostgresRepository(db)

	// Layer Business
	pService := service.NewProductService(pRepo)

	// Layer Validator (3rd party)
	v := validator.New()

	// Layer Presentation (Handler)
	pHandler := handler.NewProductHandler(pService, v) // Perhatikan: NewProductHandler perlu disesuaikan return struct-nya saja, hapus setup router di constructor lama jika ada

	// 4. Setup Router
	r := NewRouter(pHandler) // Panggil fungsi dari router.go (package main)

	// 5. Setup Server Structure for Graceful Shutdown
	srv := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: r,
	}

	// 6. Jalankan Server di background Goroutine
	go func() {
		log.Printf("🚀 Server running on http://localhost:%s", cfg.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server crash: %v", err)
		}
	}()

	// 7. Graceful Shutdown Signal Catcher
	quit := make(chan os.Signal, 1)
	// Menerima signal SIGINT (Ctrl+C) atau SIGTERM (Docker/K8s stop)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Block main goroutine sampai signal diterima
	<-quit
	log.Println("⚠️ Shutting down server gracefully...")

	// Beri batas waktu maksimal 5 detik untuk proses yang sedang berjalan agar selesai
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("✅ Server exited cleanly")
}
