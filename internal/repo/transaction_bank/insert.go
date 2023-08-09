package transaction_bank

import "github.com/phincon-backend/laza/domain/model"

func (r *TransactionBankRepo) Insert(tb model.TransactionBank) (model.TransactionBank, error) {
	tx := r.db.Create(&tb)
	err := tx.Error

	return tb, err
}
