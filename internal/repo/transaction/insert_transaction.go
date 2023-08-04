package transaction

import "github.com/phincon-backend/laza/domain/model"

func (r *TransactionRepo) Insert(ts model.Transaction) (model.Transaction, error) {
	tx := r.db.Create(&ts)
	err := tx.Error

	return ts, err
}
