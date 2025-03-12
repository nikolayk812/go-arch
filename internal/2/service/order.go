package service

import (
	"context"
	"errors"
)

func (s *service) CreateOrder(ctx context.Context, ownerID string) error {
	if ownerID == "" {
		return errors.New("owner id is empty")
	}

	if err := s.repo.CreateOrder(ctx, ownerID); err != nil {
		return err
	}

	return nil
}
