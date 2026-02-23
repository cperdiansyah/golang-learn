package service

import (
	"context"
	"fmt"
)

func (s *productService) DeleteProduct(ctx context.Context, id string) (string, error) {
	DeleteProduct, err := s.repo.Delete(ctx, id)
	if err != nil {
		return "", fmt.Errorf("Gagal delete product : %w", err)
	}
	return DeleteProduct, nil
}
