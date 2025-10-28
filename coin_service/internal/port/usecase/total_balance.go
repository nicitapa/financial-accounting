package usecase

import "context"

type TotalBalance interface {
	GetTotalBalance(ctx context.Context) (float64, error)
}
