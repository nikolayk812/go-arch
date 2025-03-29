package cart

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nikolayk812/go-arch/internal/5-screaming-tx-provider/common"
	"github.com/nikolayk812/go-arch/internal/domain"
)

type Repo interface {
	AddItem(ctx context.Context, ownerID string, item domain.CartItem) error
	Get(ctx context.Context, ownerID string) (domain.Cart, error)
	HardDelete(ctx context.Context, ownerID string) error
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

func (r repo) AddItem(ctx context.Context, ownerID string, item domain.CartItem) error {

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
