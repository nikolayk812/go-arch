package repo

import "github.com/nikolayk812/go-arch/internal/domain"

type CartToOrderFunc func(cart domain.Cart) (domain.Order, error)
