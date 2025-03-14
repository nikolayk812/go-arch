package repo

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/nikolayk812/go-arch/internal/domain"
)

func (r *repo) AddCartItem(ctx context.Context, ownerID string, item domain.CartItem) error {
	// TODO: implement

	return nil
}

func (r *repo) GetCart(ctx context.Context, ownerID string) (domain.Cart, error) {
	// TODO: implement

	return domain.Cart{}, nil
}

func (r *repo) getCart(ctx context.Context, tx pgx.Tx, ownerID string) (domain.Cart, error) {
	// TODO: implement

	return domain.Cart{}, nil
}

func (r *repo) hardDeleteCart(ctx context.Context, tx pgx.Tx, ownerID string) error {
	// TODO: implement

	return nil
}
