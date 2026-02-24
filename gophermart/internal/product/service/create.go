package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/cperdiansyah/gophermart/internal/product/entity"
)

func (s *productService) Save(ctx context.Context, req entity.CreateProductRequest) (entity.Product, error) {
	// Validation
	if req.Price <= 0 {
		return entity.Product{}, errors.New("Harga tidak boleh nol atau negatif")
	}

	if req.Stock <= 0 {
		return entity.Product{}, errors.New("stok awal tidak boleh negatif")
	}

	// Mapping dari request ke entity datbase
	product := entity.Product{
		Name:      req.Name,
		Price:     req.Price,
		Stock:     req.Stock,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Panggil repository
	if err := s.repo.Save(ctx, &product); err != nil {
		return entity.Product{}, fmt.Errorf("gagal save product : %w", err)

	}
	return product, nil
}
