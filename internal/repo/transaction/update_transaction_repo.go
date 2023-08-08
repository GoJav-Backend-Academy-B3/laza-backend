package transaction

import (
	"errors"
	"github.com/phincon-backend/laza/domain/model"
)

func (r *TransactionRepo) Update(id any, ts model.Transaction) (model.Transaction, error) {
	if id.(string) != ts.Id {
		return ts, errors.New("id not found")
	}

	tx := r.db.Save(&ts)
	err := tx.Error

	return ts, err
}
