package dbstore

import (
	"coin_service/internal/errs"
	"database/sql"
	"errors"
)

func (t *TransactionStorage) translateError(err error) error {
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return errs.ErrNotfound
	default:
		return err
	}
}
