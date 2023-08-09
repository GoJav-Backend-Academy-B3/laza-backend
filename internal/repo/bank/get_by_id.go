package bank

import (
	"github.com/phincon-backend/laza/domain/model"
)

func (r *BankRepo) GetById(id any) (e model.Bank, err error) {
	tx := r.db.First(&e, "id = ?", id)
	err = tx.Error
	return
}
