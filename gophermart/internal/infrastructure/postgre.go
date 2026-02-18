package infrastructure

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib" // Driver tetap di sini
)

func NewPostgresDB(connectionString string) (*sql.DB, error) {
	db, err := sql.Open("pgx", connectionString)
	if err != nil {
		return nil, fmt.Errorf("gagal membuka driver: %w", err)
	}

	// Tes koneksi (Fail Fast)
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("gagal connect ke db: %w", err)
	}

	return db, nil
}
