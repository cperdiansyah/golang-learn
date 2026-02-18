package entity

import "time"

type Product struct {
	ID        int64     `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Price     float64   `json:"price" db:"price"`
	Stock     int       `json:"stock" db:"stock"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateProductRequest struct {
	// required: wajib isi, min=3: minimal 3 karakter
	Name string `json:"name" validate:"required,min=3"`

	// gt=0: greater than 0 (harus positif)
	Price float64 `json:"price" validate:"required,gt=0"`

	// gte=0: greater than or equal 0 (boleh 0, gak boleh minus)
	Stock int `json:"stock" validate:"gte=0"`
}
