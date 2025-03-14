package order

import (
	"context"
	"github.com/nikolayk812/go-arch/internal/4-hybrid/repo"
	"github.com/nikolayk812/go-arch/internal/domain"
)

type Service interface {
	CreateOrder(ctx context.Context, ownerID string) error
}

type Repo interface {
	CreateOrder(ctx context.Context, ownerID string, cartToOrder repo.CartToOrderFunc) error
}

type service struct {
	repo Repo
}

func NewService(repo Repo) Service {
	return &service{
		repo: repo,
	}
}

func (s service) CreateOrder(ctx context.Context, ownerID string) error {
	return s.repo.CreateOrder(ctx, ownerID, cartToOrder)
}

func cartToOrder(cart domain.Cart) (domain.Order, error) {
	//TODO implement

	return domain.Order{}, nil
}
