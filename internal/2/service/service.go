package service

import (
	"context"
	"github.com/nikolayk812/go-arch/internal/2/repo"
	"github.com/nikolayk812/go-arch/internal/domain"
)

type Service interface {
	// cart section
	AddCartItem(ctx context.Context, ownerID string, item domain.CartItem) error

	// order section
	CreateOrder(ctx context.Context, ownerID string) error
}

type service struct {
	repo repo.Repo
}

func New(repo repo.Repo) Service {
	return &service{repo: repo}
}
