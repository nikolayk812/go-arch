package cart

import (
	"context"
	"github.com/nikolayk812/go-arch/internal/domain"
)

type Service interface {
	AddItem(ctx context.Context, ownerID string, item domain.CartItem) error
	Get(ctx context.Context, ownerID string) (domain.Cart, error)
}

type service struct {
	repo Repo
}

func NewService(repo Repo) Service {
	return &service{repo: repo}
}

func (s service) AddItem(ctx context.Context, ownerID string, item domain.CartItem) error {
	// TODO: business validation logic

	return s.repo.AddItem(ctx, ownerID, item)
}

func (s service) Get(ctx context.Context, ownerID string) (domain.Cart, error) {
	return s.repo.Get(ctx, ownerID)
}
