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
	repo       Repo
	txProvider TransactionProvider
}

func NewService(repo Repo, txProvider TransactionProvider) Service {
	return &service{
		repo:       repo,
		txProvider: txProvider,
	}
}

func (s service) CreateOrder(ctx context.Context, ownerID string) error {
	err := s.txProvider.Transact(ctx, func(adapters Adapters) error {
		cart, err := adapters.cartRepo.Get(ctx, ownerID)
		if err != nil {
			return err
		}

		order, err := cartToOrder(cart)
		if err != nil {
			return err
		}

		err = adapters.repo.Create(ctx, ownerID, order)
		if err != nil {
			return err
		}

		err = adapters.cartRepo.HardDelete(ctx, ownerID)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func cartToOrder(cart domain.Cart) (domain.Order, error) {
	//TODO implement

	return domain.Order{}, nil
}
