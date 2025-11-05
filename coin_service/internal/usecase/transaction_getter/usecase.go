package transaction_getter

import (
	"coin_service/internal/config"
	"coin_service/internal/domain"
	"coin_service/internal/errs"
	"coin_service/internal/port/driven"
	"context"
	"errors"
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

func (u *UseCase) GetTransactionByID(ctx context.Context, id int) (transaction domain.Transaction, err error) {
	var tx domain.Transaction

	tx, err = u.transactionStorage.GetTransactionByID(ctx, id)
	if err != nil {
		if errors.Is(err, errs.ErrNotfound) {
			return domain.Transaction{}, errs.ErrTransactionNotFound
		}
		return domain.Transaction{}, err
	}
	return tx, nil
}
