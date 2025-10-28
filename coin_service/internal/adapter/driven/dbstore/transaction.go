package dbstore

import (
	"coin_service/internal/domain"
	"context"
	"time"

	"github.com/jmoiron/sqlx"
)

type TransactionStorage struct {
	db *sqlx.DB
}

func NewTransactionStorage(db *sqlx.DB) *TransactionStorage {
	return &TransactionStorage{db: db}
}

type Transaction struct {
	ID          int             `db:"id"`
	Category    domain.Category `db:"category"`
	Amount      float64         `db:"amount"`
	Currency    string          `db:"currency"`
	Timestamp   time.Time       `db:"timestamp"`
	Description string          `db:"description"`
}

func (t *Transaction) ToDomain() *domain.Transaction {
	return &domain.Transaction{
		ID:          t.ID,
		Category:    t.Category,
		Amount:      t.Amount,
		Currency:    t.Currency,
		Timestamp:   t.Timestamp,
		Description: t.Description,
	}
}

func (t *Transaction) FromDomain(d domain.Transaction) {
	t.ID = d.ID
	t.Category = d.Category
	t.Amount = d.Amount
	t.Currency = d.Currency
	t.Timestamp = d.Timestamp
	t.Description = d.Description
}

func (t *TransactionStorage) CreateTransaction(ctx context.Context, transaction domain.Transaction) (err error) {
	var dbTransaction Transaction
	dbTransaction.FromDomain(transaction)
	_, err = t.db.ExecContext(ctx,
		`INSERT INTO transactions (category, amount, currency, description)
	 VALUES ($1, $2, $3, $4);`,
		dbTransaction.Category,
		dbTransaction.Amount,
		dbTransaction.Currency,
		dbTransaction.Description,
	)
	if err != nil {
		return t.translateError(err)
	}
	return nil
}

func (t *TransactionStorage) GetAllTransactions(ctx context.Context) ([]domain.Transaction, error) {
	var txs []domain.Transaction
	err := t.db.SelectContext(ctx, &txs, `
		SELECT id, category, amount, currency, timestamp, description
		FROM transactions
		ORDER BY timestamp DESC
	`)
	return txs, err
}

func (t *TransactionStorage) GetTransactionByID(ctx context.Context, id int) (transaction domain.Transaction, err error) {
	var dbTransaction Transaction
	if err := t.db.GetContext(ctx, &dbTransaction, `
	SELECT id, category, amount, currency, timestamp, description 
	FROM transactions
	WHERE id = $1`, id); err != nil {
		return domain.Transaction{}, t.translateError(err)
	}
	return *dbTransaction.ToDomain(), nil
}

func (t *TransactionStorage) GetTotalBalance(ctx context.Context) (float64, error) {
	var total float64

	query := `
		SELECT
			COALESCE(SUM(
				CASE
					WHEN category = 'INCOME' THEN amount
					WHEN category = 'EXPENSE' THEN -amount
					ELSE 0
				END
			), 0) AS total
		FROM transactions;
	`

	err := t.db.GetContext(ctx, &total, query)
	if err != nil {
		return 0, t.translateError(err)
	}

	return total, nil
}
