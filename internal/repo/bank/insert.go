package bank

import (
	"github.com/phincon-backend/laza/domain/model"
)

func (r *BankRepo) Insert(dao model.Bank) (e model.Bank, err error) {
	tx := r.db.Create(&dao).Scan(&e)
	err = tx.Error
	return
}
