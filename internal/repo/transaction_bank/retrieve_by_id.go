package transaction_bank

import "github.com/phincon-backend/laza/domain/model"

func (r *TransactionBankRepo) GetById(id any) (tb model.TransactionBank, err error) {
	tx := r.db.First(&tb, id)
	err = tx.Error
	return
}
