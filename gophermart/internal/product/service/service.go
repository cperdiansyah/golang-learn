package service

import (
	"context"

	"github.com/cperdiansyah/gophermart/internal/product/entity"
	"github.com/cperdiansyah/gophermart/internal/product/repository"
)

type Service interface {
	CreateProduct(ctx context.Context, req entity.CreateProductRequest) (entity.Product, error)
	GetAllProduct(ctx context.Context) ([]entity.Product, error)
	GetProductByID(ctx context.Context, id string) (*entity.Product, error)
	UpdateProduct(ctx context.Context, req entity.CreateProductRequest, id string) (entity.Product, error)
	DeleteProduct(ctx context.Context, id string) (string, error)
}

type productService struct {
	repo repository.Repository
}

func NewProductService(repo repository.Repository) Service {
	return &productService{repo: repo}
}
