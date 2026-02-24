package service

import (
	"context"
	"fmt"

	"github.com/cperdiansyah/gophermart/internal/modules/product/entity"
)

func (s *productService) FindAll(ctx context.Context, limit int, offset int) ([]entity.Product, error) {

	products, err := s.repo.FindAll(ctx, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("gagal get all product : %w", err)
	}
	return products, nil
}

func (s *productService) FindByID(ctx context.Context, id string) (*entity.Product, error) {
	product, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("gagal get product by id : %w", err)
	}
	return product, nil
}
