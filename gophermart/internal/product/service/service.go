package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/cperdiansyah/gophermart/internal/product/entity"
	"github.com/cperdiansyah/gophermart/internal/product/repository"
)

type Service interface {
	CreateProduct(ctx context.Context, req entity.CreateProductRequest) (int, error)
}

type productService struct {
	repo repository.Repository
}

func NewProductService(repo repository.Repository) Service {
	return &productService{repo: repo}
}

func (s *productService) CreateProduct(ctx context.Context, req entity.CreateProductRequest) (int, error) {
	// Validation
	if req.Price <= 0 {
		return 0, errors.New("Harga tidak boleh nol atau negatif")
	}

	if req.Stock <= 0 {
		return 0, errors.New("stok awal tidak boleh negatif")
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
		return 0, fmt.Errorf("gagal save product : %w", err)
	}
	return int(product.ID), nil
}
