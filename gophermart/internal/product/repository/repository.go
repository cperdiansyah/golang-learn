package repository

import (
	"context"

	"github.com/cperdiansyah/gophermart/internal/product/entity"
)

type Repository interface {
	Save(ctx context.Context, p *entity.Product) error
	FindAll(ctx context.Context) ([]entity.Product, error)
	FindByID(ctx context.Context, id int) (*entity.Product, error)
	Update(ctx context.Context, p *entity.Product) error
	Delete(ctx context.Context, id int) error
}
