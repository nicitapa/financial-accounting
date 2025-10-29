package usecase

import (
	"coin_service/internal/domain"
	"context"
)

type TransactionGetter interface {
	GetTransactionByID(ctx context.Context, id int) (transaction domain.Transaction, err error)
}
