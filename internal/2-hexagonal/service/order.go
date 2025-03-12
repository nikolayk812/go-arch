package service

import (
	"context"
	"errors"
	"github.com/nikolayk812/go-arch/internal/domain"
)

func (s *service) CreateOrder(ctx context.Context, ownerID string) error {
	if ownerID == "" {
		return errors.New("owner id is empty")
	}

	if err := s.repo.CreateOrder(ctx, ownerID, cartToOrder); err != nil {
		return err
	}

	return nil
}

func cartToOrder(cart domain.Cart) (domain.Order, error) {
	// TODO: implement

	return domain.Order{}, nil
}
