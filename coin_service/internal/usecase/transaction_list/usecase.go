package transaction_list

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

func (u *UseCase) GetAllTransactions(ctx context.Context) ([]domain.Transaction, error) {
	transactions, err := u.transactionStorage.GetAllTransactions(ctx)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}
