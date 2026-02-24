package repository

import (
	"context"

	"github.com/cperdiansyah/gophermart/internal/product/entity"
)

type Repository interface {
	Save(ctx context.Context, p *entity.Product) error
	FindAll(ctx context.Context, limit int, offset int) ([]entity.Product, error)
	FindByID(ctx context.Context, id string) (*entity.Product, error)
	Update(ctx context.Context, p *entity.Product) (entity.Product, error)
	Delete(ctx context.Context, id string) (string, error)
}
