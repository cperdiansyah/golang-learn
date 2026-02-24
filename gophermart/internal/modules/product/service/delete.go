package service

import (
	"context"
	"fmt"

	"github.com/cperdiansyah/gophermart/internal/modules/product/entity"
)

func (s *productService) Delete(ctx context.Context, id string) (string, error) {
	product, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return "", fmt.Errorf("gagal cek product: %w", err)
	}
	if product == nil {
		return "", entity.ErrNotFound
	}

	DeleteProduct, err := s.repo.Delete(ctx, id)
	if err != nil {
		return "", fmt.Errorf("Gagal delete product : %w", err)
	}
	return DeleteProduct, nil
}
