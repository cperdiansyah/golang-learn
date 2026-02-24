package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	DBConnStr string
	Port      string
}

func LoadConfig() AppConfig {
	// Memuat nilai dari file .env (jika tidak ada akan di-skip ke global env vars system)
	err := godotenv.Load()
	if err != nil {
		log.Println("Note: .env file not found, loading from global environment variables")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port
	}

	dbConnStr := os.Getenv("DB_CONN_STR")
	if dbConnStr == "" {
		log.Fatal("CRITICAL: DB_CONN_STR environment variable is required")
	}

	return AppConfig{
		DBConnStr: dbConnStr,
		Port:      port,
	}
}
