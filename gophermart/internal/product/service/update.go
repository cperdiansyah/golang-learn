package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/cperdiansyah/gophermart/internal/product/entity"
)

func (s *productService) UpdateProduct(ctx context.Context, req entity.CreateProductRequest, id string) (entity.Product, error) {
	if id == "" {
		return entity.Product{}, errors.New("ID tidak boleh kosong")
	}
	// Validation
	if req.Price <= 0 {
		return entity.Product{}, errors.New("Harga tidak boleh nol atau negatif")
	}

	if req.Stock <= 0 {
		return entity.Product{}, errors.New("stok awal tidak boleh negatif")
	}

	product, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return entity.Product{}, fmt.Errorf("gagal update product : %w", err)
	}
	Product := entity.Product{
		ID:        product.ID,	
		Name:      product.Name,
		Price:     product.Price,
		Stock:     product.Stock,
		CreatedAt: product.CreatedAt,
		UpdatedAt: time.Now(),
	}

	updatedProduct, err := s.repo.Update(ctx, &Product)
	if err != nil {
		return entity.Product{}, fmt.Errorf("Gagal update product : %w", err)
	}
	return updatedProduct, nil

}
