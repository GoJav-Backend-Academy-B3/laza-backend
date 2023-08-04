package transaction

import "github.com/phincon-backend/laza/domain/model"

func (r *TransactionRepo) GetAll() (transaction []model.Transaction, err error) {
	tx := r.db.Find(&transaction)
	err = tx.Error
	return
}
