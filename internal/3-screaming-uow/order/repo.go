package order

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nikolayk812/go-arch/internal/domain"
)

type Repo interface {
	Create(ctx context.Context, tx pgx.Tx, ownerID string, order domain.Order) error
}

type repo struct {
	pool *pgxpool.Pool
}

func NewRepo(pool *pgxpool.Pool) Repo {
	return &repo{pool: pool}
}

func (r repo) Create(ctx context.Context, tx pgx.Tx, ownerID string, order domain.Order) error {
	//TODO implement insertion

	return nil
}
