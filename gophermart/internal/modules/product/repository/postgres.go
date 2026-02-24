package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"

	"github.com/cperdiansyah/gophermart/internal/modules/product/entity"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) Repository {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) Save(ctx context.Context, p *entity.Product) error {
	query := `INSERT INTO products (name, price, stock,  created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id`

	//QueryRowContext dipake kalau kita mau dapet balikan data (ID yang baru kebuat)
	//QueryRowContext mengembalikan 1 baris data
	//QueryContext mengembalikan banyak baris data
	err := r.db.QueryRowContext(ctx, query, p.Name, p.Price, p.Stock, p.CreatedAt, p.UpdatedAt).Scan(&p.ID)
	if err != nil {
		return fmt.Errorf("gagal insert product: %w", err)
	}
	return nil
}

func (r *PostgresRepository) FindAll(ctx context.Context, limit int, offset int) ([]entity.Product, error) {
	query := `
	SELECT id, name, price, stock, created_at, updated_at, deleted_at
	From products
	WHERE deleted_at IS NULL
	ORDER BY created_at DESC
	LIMIT $1 OFFSET $2
	`
	rows, err := r.db.QueryContext(ctx, query, limit, offset)

	if err != nil {
		return nil, fmt.Errorf("Gagal query find all : %w", err)
	}

	defer rows.Close()

	var products []entity.Product
	for rows.Next() {
		var p entity.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Stock, &p.CreatedAt, &p.UpdatedAt, &p.DeletedAt); err != nil {
			return nil, fmt.Errorf("Gagal scan product : %w", err)
		}
		products = append(products, p)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Gagal query find all : %w", err)

	}
	return products, nil
}

func (r *PostgresRepository) FindByID(ctx context.Context, id string) (*entity.Product, error) {
	query := `
	SELECT id, name, price, stock, created_at, updated_at, deleted_at
	FROM products
	WHERE id = $1 AND deleted_at IS NULL
	 `
	var p entity.Product
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return nil, fmt.Errorf("invalid ID format: %w", err)
	}
	err = r.db.QueryRowContext(ctx, query, idInt).Scan(
		&p.ID,
		&p.Name,
		&p.Price,
		&p.Stock,
		&p.CreatedAt,
		&p.UpdatedAt,
		&p.DeletedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil // akan return custom error not found
	}
	if err != nil {
		return nil, fmt.Errorf("Gagal query find by id : %w", err)
	}
	return &p, nil
}

func (r *PostgresRepository) Update(ctx context.Context, p *entity.Product) (entity.Product, error) {
	query := `
	UPDATE products
	SET name=$1, price=$2, stock=$3, updated_at=$4
	WHERE id=$5
	`
	result, err := r.db.ExecContext(ctx, query, p.Name, p.Price, p.Stock, p.UpdatedAt, p.ID)
	if err != nil {
		return entity.Product{}, fmt.Errorf("gagal update : %w", err)
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return entity.Product{}, fmt.Errorf("Product dengan id %d tidak ditemukan", p.ID)
	}
	return *p, nil
}

func (r *PostgresRepository) Delete(ctx context.Context, id string) (string, error) {
	query := `
		UPDATE products
		SET deleted_at = NOW()
		WHERE id = $1
	`
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return "", fmt.Errorf("invalid ID format: %w", err)
	}
	_, err = r.db.ExecContext(ctx, query, idInt)
	if err != nil {
		return "", fmt.Errorf("Gagal delete:%w", err)
	}
	return id, nil
}
