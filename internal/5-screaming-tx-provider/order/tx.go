package order

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nikolayk812/go-arch/internal/5-screaming-tx-provider/cart"
	"github.com/nikolayk812/go-arch/internal/5-screaming-tx-provider/common"
	"github.com/nikolayk812/go-arch/internal/domain"
)

type TransactionProvider struct {
	pool *pgxpool.Pool
}

type Adapters struct {
	repo     Repo
	cartRepo cartRepo
}

type cartRepo interface {
	Get(ctx context.Context, ownerID string) (domain.Cart, error)
	HardDelete(ctx context.Context, ownerID string) error
}

func NewTransactionProvider(pool *pgxpool.Pool) *TransactionProvider {
	return &TransactionProvider{
		pool: pool,
	}
}

func (p *TransactionProvider) Transact(ctx context.Context, txFunc func(adapters Adapters) error) error {
	return common.RunInTx(ctx, p.pool, func(tx pgx.Tx) error {
		adapters := Adapters{
			repo:     NewRepoWithTx(tx),
			cartRepo: cart.NewRepoWithTx(tx),
		}

		return txFunc(adapters)
	})
}
