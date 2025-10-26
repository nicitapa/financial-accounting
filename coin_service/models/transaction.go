package models

import "time"

type Transaction struct {
	ID          uint      `json:"ID" db:"id"`
	UserID      string    `json:"UserID" db:"user_id"`
	Category    string    `json:"Category" db:"category"`
	Amount      float64   `json:"Amount" db:"amount"`
	Currency    string    `json:"currency" db:"currency"`
	Timestamp   time.Time `json:"timestamp" db:"timestamp"`
	Description string    `json:"description" db:"description"`
}
