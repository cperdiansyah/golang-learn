package service

import (
	"context"

	"github.com/cperdiansyah/gophermart/internal/product/entity"
	"github.com/cperdiansyah/gophermart/internal/product/repository"
)

type Service interface {
	Save(ctx context.Context, req entity.CreateProductRequest) (entity.Product, error)
	FindAll(ctx context.Context, limit int, offset int) ([]entity.Product, error)
	FindByID(ctx context.Context, id string) (*entity.Product, error)
	Update(ctx context.Context, req entity.UpdateProductRequest, id string) (entity.Product, error)
	Delete(ctx context.Context, id string) (string, error)
}

type productService struct {
	repo repository.Repository
}

func NewProductService(repo repository.Repository) Service {
	return &productService{repo: repo}
}
