package usecase

import (
	"coin_service/internal/adapter/driven/dbstore"
	"coin_service/internal/config"
	"coin_service/internal/port/usecase"
	totalbalance "coin_service/internal/usecase/total_balance"
	transactioncreator "coin_service/internal/usecase/transaction_creator"
	transactiongetter "coin_service/internal/usecase/transaction_getter"
	transactionlist "coin_service/internal/usecase/transaction_list"
)

type UseCases struct {
	TransactionCreator usecase.TransactionCreator
	TransactionGetter  usecase.TransactionGetter
	TransactionList    usecase.TransactionList
	TotalBalance       usecase.TotalBalance
}

func New(cfg *config.Config, store *dbstore.DBStore) *UseCases {
	return &UseCases{
		TransactionCreator: transactioncreator.New(cfg, store.TransactionStorage),
		TransactionGetter:  transactiongetter.New(cfg, store.TransactionStorage),
		TransactionList:    transactionlist.New(cfg, store.TransactionStorage),
		TotalBalance:       totalbalance.New(cfg, store.TransactionStorage),
	}
}
