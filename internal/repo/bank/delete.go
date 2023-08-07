package bank

import "github.com/phincon-backend/laza/domain/model"

func (r *BankRepo) Delete(id any) (err error) {
	var data model.Bank
	tx := r.db.Delete(data, "id = ?", id)
	err = tx.Error
	return
}
