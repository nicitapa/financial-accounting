package dbstore

import "github.com/jmoiron/sqlx"

type DBStore struct {
	TransactionStorage *TransactionStorage
}

func New(db *sqlx.DB) *DBStore {
	return &DBStore{
		TransactionStorage: NewTransactionStorage(db),
	}
}
