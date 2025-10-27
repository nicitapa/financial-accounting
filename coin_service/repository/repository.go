package repository

import (
	"coin_service/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type TransactionRepository struct {
	DB *sqlx.DB
}

func NewTransactionRepository(db *sqlx.DB) *TransactionRepository {
	return &TransactionRepository{DB: db}
}

func (r *TransactionRepository) CreateTransactions(tx models.Transaction) error {
	_, err := r.DB.NamedExec(`INSERT INTO transactions
	(user_id, type, category, amount, currency, timestamp, description)
	VALUES (:user_id, :type, :category, :amount, :currency, :timestamp, :description)`, tx)
	return err
}

func (r *TransactionRepository) GetAllTransactions() ([]models.Transaction, error) {
	var txs []models.Transaction
	err := r.DB.Select(&txs, `
		SELECT id, user_id, type, category, amount, currency, timestamp, description
		FROM transactions
		ORDER BY timestamp DESC
	`)
	return txs, err
}
