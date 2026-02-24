package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/cperdiansyah/gophermart/internal/modules/product/entity"
)

func (s *productService) Update(ctx context.Context, req entity.UpdateProductRequest, id string) (entity.Product, error) {
	if id == "" {
		return entity.Product{}, errors.New("ID tidak boleh kosong")
	}

	product, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return entity.Product{}, fmt.Errorf("gagal update product : %w", err)
	}

	// Tangani kasus Product not found (termasuk kalau sudah di soft-delete)
	if product == nil {
		return entity.Product{}, entity.ErrNotFound
	}

	// Lakukan map nilai baru jika field diprovide (pointer isn't nil).
	// Jika nil, kita keep nilai lama dari database.
	Product := entity.Product{
		ID:        product.ID,
		Name:      product.Name,
		Price:     product.Price,
		Stock:     product.Stock,
		CreatedAt: product.CreatedAt,
		UpdatedAt: time.Now(),
	}

	if req.Name != nil {
		Product.Name = *req.Name
	}
	if req.Price != nil {
		Product.Price = *req.Price
	}
	if req.Stock != nil {
		Product.Stock = *req.Stock
	}

	updatedProduct, err := s.repo.Update(ctx, &Product)
	if err != nil {
		return entity.Product{}, fmt.Errorf("Gagal update product : %w", err)
	}
	return updatedProduct, nil

}
