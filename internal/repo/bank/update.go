package bank

import (
	"github.com/phincon-backend/laza/domain/model"
)

func (r *BankRepo) Update(id any, dao model.Bank) (e model.Bank, err error) {
	tx := r.db.Model(dao).Where("id = ?", id).Updates(&dao).Scan(&e)
	err = tx.Error
	return
}
