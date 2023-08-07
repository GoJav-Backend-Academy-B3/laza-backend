package bank

import (
	"github.com/phincon-backend/laza/domain/model"
)

func (r *BankRepo) GetAll() (es []model.Bank, err error) {
	tx := r.db.Find(&es).Scan(&es)
	err = tx.Error
	return
}
