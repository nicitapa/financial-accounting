package usecase

import (
	"coin_service/internal/usecase/total_balance"
	"context"
)

type TotalBalance interface {
	GetTotalBalance(ctx context.Context) (*total_balance.BalanceResponse, error)
}
