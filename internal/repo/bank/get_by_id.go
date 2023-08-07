package bank

import (
	"github.com/phincon-backend/laza/domain/model"
)

func (r *BankRepo) GetBankById(id any) (e model.Bank, err error) {
	tx := r.db.First(&e, "id = ?", id)
	err = tx.Error
	return
}
