package domain

import "time"

type Category string

const (
	Income  Category = "INCOME"
	Expense Category = "EXPENSE"
)

type Transaction struct {
	ID          int
	Category    Category
	Amount      float64
	Currency    string
	Timestamp   time.Time
	Description string
}
