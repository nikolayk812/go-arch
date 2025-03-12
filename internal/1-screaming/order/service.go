package order

import (
	"context"
	"github.com/nikolayk812/go-arch/internal/domain"
)

type Service interface {
	CreateOrder(ctx context.Context, ownerID string) error
}

type CartRepo interface {
	Get(ctx context.Context, ownerID string) (domain.Cart, error)
	HardDelete(ctx context.Context, ownerID string) error
}

type service struct {
	repo     Repo
	cartRepo CartRepo
}

func NewService(repo Repo, cartRepo CartRepo) Service {
	return &service{
		repo:     repo,
		cartRepo: cartRepo,
	}
}

func (s service) CreateOrder(ctx context.Context, ownerID string) error {
	// tx1
	cart, err := s.cartRepo.Get(ctx, ownerID)
	if err != nil {
		return err
	}

	order, err := cartToOrder(cart)
	if err != nil {
		return err
	}

	// tx2
	if err := s.repo.Create(ctx, order); err != nil {
		return err
	}

	// tx3
	if err := s.cartRepo.HardDelete(ctx, ownerID); err != nil {
		return err
	}

	return nil
}

func cartToOrder(cart domain.Cart) (domain.Order, error) {
	//TODO implement

	return domain.Order{}, nil
}
