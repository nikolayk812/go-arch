package service

import (
	"context"
	"github.com/nikolayk812/go-arch/internal/domain"
)

func (s *service) AddCartItem(ctx context.Context, ownerID string, item domain.CartItem) error {
	return s.repo.AddCartItem(ctx, ownerID, item)
}
