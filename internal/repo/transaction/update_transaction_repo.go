package transaction

import "github.com/phincon-backend/laza/domain/model"

func (r *TransactionRepo) Update(ts model.Transaction) (model.Transaction, error) {
	tx := r.db.Save(&ts)
	err := tx.Error

	return ts, err
}
