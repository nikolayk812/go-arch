package cart

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nikolayk812/go-arch/internal/domain"
)

type Repo interface {
	AddItem(ctx context.Context, ownerID string, item domain.CartItem) error
	Get(ctx context.Context, ownerID string) (domain.Cart, error)
	HardDelete(ctx context.Context, ownerID string) error
}

type repo struct {
	pool *pgxpool.Pool
}

func NewRepo(pool *pgxpool.Pool) Repo {
	return &repo{pool: pool}
}

func (r repo) AddItem(ctx context.Context, ownerID string, item domain.CartItem) error {
	//TODO implement

	return nil
}

func (r repo) Get(ctx context.Context, ownerID string) (domain.Cart, error) {
	//TODO implement

	return domain.Cart{}, nil
}

func (r repo) HardDelete(ctx context.Context, ownerID string) error {
	//TODO implement

	return nil
}
