package transaction

import "github.com/phincon-backend/laza/domain/model"

func (r *TransactionRepo) Get(id uint64) (transaction model.Transaction, err error) {
	tx := r.db.First(&transaction, id)
	err = tx.Error
	return
}
