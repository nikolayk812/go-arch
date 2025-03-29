package order

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nikolayk812/go-arch/internal/5-screaming-tx-provider/common"
	"github.com/nikolayk812/go-arch/internal/domain"
)

type Repo interface {
	Create(ctx context.Context, ownerID string, order domain.Order) error
}

type repo struct {
	db common.DB
}

func NewRepo(pool *pgxpool.Pool) Repo {
	return &repo{db: pool}
}

func NewRepoWithTx(tx pgx.Tx) Repo {
	return &repo{db: tx}
}

func (r repo) Create(ctx context.Context, ownerID string, order domain.Order) error {
	//TODO implement insertion

	return nil
}
