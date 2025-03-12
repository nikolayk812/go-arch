package repo

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nikolayk812/go-arch/internal/domain"
)

type Repo interface {
	// cart section
	AddCartItem(ctx context.Context, ownerID string, item domain.CartItem) error

	// order section
	CreateOrder(ctx context.Context, ownerID string) error
}

type repo struct {
	pool        *pgxpool.Pool
	cartToOrder CartToOrderFunc // could be a param in CreateOrder instead
}

func New(pool *pgxpool.Pool, cartToOrder CartToOrderFunc) Repo {
	return &repo{
		pool:        pool,
		cartToOrder: cartToOrder,
	}
}
