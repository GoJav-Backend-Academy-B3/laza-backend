package bank

import (
	"github.com/phincon-backend/laza/domain/model"
)

func (r *BankRepo) ExistsBank(name string) bool {
	var data model.Bank
	err := r.db.
		First(&data, "bank_name = ?", name)
	return err.Error == nil
}
