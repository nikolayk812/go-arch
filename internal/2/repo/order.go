package repo

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/nikolayk812/go-arch/internal/domain"
)

func (r *repo) CreateOrder(ctx context.Context, ownerID string) error {
	// for brevity, not production-ready tx management
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Commit(ctx)

	cart, err := r.getCart(ctx, tx, ownerID)
	if err != nil {
		return err
	}

	order, err := r.cartToOrder(cart)
	if err != nil {
		return err
	}

	if err := r.insertOrder(ctx, tx, ownerID, order); err != nil {
		return err
	}

	if err := r.hardDeleteCart(ctx, tx, ownerID); err != nil {
		return err
	}

	// TODO: to be super-safe verify number deleted cart items is same (expected) as in cart read

	return nil
}

func (r *repo) insertOrder(ctx context.Context, tx pgx.Tx, ownerID string, order domain.Order) error {
	// TODO: implement

	return nil
}
