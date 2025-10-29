package driven

import (
	"coin_service/internal/domain"
	"context"
)

type TransactionStorage interface {
	CreateTransaction(ctx context.Context, transaction domain.Transaction) (err error)
	GetAllTransactions(ctx context.Context) ([]domain.Transaction, error)
	GetTransactionByID(ctx context.Context, id int) (transaction domain.Transaction, err error)
	GetTotalBalance(ctx context.Context) (float64, error)
}
