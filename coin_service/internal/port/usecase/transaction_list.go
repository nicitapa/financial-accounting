package usecase

import (
	"coin_service/internal/domain"
	"context"
)

type TransactionList interface {
	GetAllTransactions(ctx context.Context) ([]domain.Transaction, error)
}
