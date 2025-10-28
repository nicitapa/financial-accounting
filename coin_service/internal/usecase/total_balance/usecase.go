package total_balance

import (
	"coin_service/internal/config"
	"coin_service/internal/port/driven"
	"context"
	"fmt"
	"log"
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

func (u *UseCase) GetTotalBalance(ctx context.Context) (float64, error) {
	total, err := u.transactionStorage.GetTotalBalance(ctx)
	if err != nil {
		log.Println("failed to get balance:", err)
		return 0, nil
	}

	if total >= 0 {
		fmt.Printf("💰 Прибыль: %.2f\n", total)
	} else {
		fmt.Printf("🔻 Убыток: %.2f\n", total)
	}
	return 0, nil
}
