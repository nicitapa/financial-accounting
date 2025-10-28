package total_balance

import (
	"coin_service/internal/config"
	"coin_service/internal/port/driven"
	"context"
	"log"
)

type UseCase struct {
	cfg                *config.Config
	transactionStorage driven.TransactionStorage
}
type BalanceResponse struct {
	Total   float64 `json:"total"`
	Status  string  `json:"status"`
	Message string  `json:"message,omitempty"`
}

func New(cfg *config.Config, transactionStorage driven.TransactionStorage) *UseCase {
	return &UseCase{
		cfg:                cfg,
		transactionStorage: transactionStorage,
	}
}

func (u *UseCase) GetTotalBalance(ctx context.Context) (*BalanceResponse, error) {
	total, err := u.transactionStorage.GetTotalBalance(ctx)
	if err != nil {
		log.Println("failed to get balance:", err)
		return &BalanceResponse{
			Total:   0,
			Status:  "error",
			Message: err.Error(),
		}, err
	}

	status := `"Бабки" не проблема💰`
	if total < 0 {
		status = "Дружок, уже должок🔻"
	} else if total == 0 {
		status = `Денюжки "тютю🕳️"`
	}

	return &BalanceResponse{
		Total:  total,
		Status: status,
	}, nil
}
