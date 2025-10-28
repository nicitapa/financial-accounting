package usecase

import (
	"coin_service/internal/domain"
	"context"
)

type TransactionCreator interface {
	CreateTransaction(ctx context.Context, transaction domain.Transaction) error
}
