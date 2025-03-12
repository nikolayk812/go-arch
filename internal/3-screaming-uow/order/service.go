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
	repo Repo
	uow  UnitOfWork
}

func NewService(repo Repo, uow UnitOfWork) Service {
	return &service{
		repo: repo,
		uow:  uow,
	}
}

func (s service) CreateOrder(ctx context.Context, ownerID string) error {
	if err := s.uow.CreateOrder(ctx, ownerID); err != nil {
		return err
	}

	return nil
}

func cartToOrder(cart domain.Cart) (domain.Order, error) {
	//TODO implement

	return domain.Order{}, nil
}
