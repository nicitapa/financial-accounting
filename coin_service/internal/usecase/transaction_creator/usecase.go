package transaction_creator

import (
	"coin_service/internal/config"
	"coin_service/internal/domain"
	"coin_service/internal/port/driven"
	"context"
)

type UseCase struct {
	cfg                *config.Config
	transactionStorage driven.TransactionStorage
}

func New(cfg *config.Config, transactionStorage driven.TransactionStorage) *UseCase {
	return &UseCase{
		cfg:                cfg,
		transactionStorage: transactionStorage,
	}
}

func (u *UseCase) CreateTransaction(ctx context.Context, transaction domain.Transaction) error {
	err := u.transactionStorage.CreateTransaction(ctx, transaction)
	if err != nil {
		return err
	}
	return nil
}
