package order

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nikolayk812/go-arch/internal/domain"
)

type UnitOfWork interface {
	CreateOrder(ctx context.Context, ownerID string) error
}

type cartRepo interface {
	GetTx(ctx context.Context, tx pgx.Tx, ownerID string) (domain.Cart, error)
	HardDelete(ctx context.Context, tx pgx.Tx, ownerID string) error
}

type uow struct {
	pool     *pgxpool.Pool
	repo     Repo
	cartRepo cartRepo
}

func NewUnitOfWork(repo Repo, cartRepo cartRepo) UnitOfWork {
	return &uow{
		repo:     repo,
		cartRepo: cartRepo,
	}
}

func (u uow) CreateOrder(ctx context.Context, ownerID string) error {
	// for brevity, not production-ready tx management
	tx, err := u.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Commit(ctx)

	// tx1
	cart, err := u.cartRepo.GetTx(ctx, tx, ownerID)
	if err != nil {
		return err
	}

	order, err := cartToOrder(cart)
	if err != nil {
		return err
	}

	if err := u.repo.Create(ctx, tx, ownerID, order); err != nil {
		return err
	}

	if err := u.cartRepo.HardDelete(ctx, tx, ownerID); err != nil {
		return err
	}

	return nil
}
