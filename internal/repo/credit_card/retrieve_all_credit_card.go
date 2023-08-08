package creditcard

import "github.com/phincon-backend/laza/domain/model"

func (r *CreditCardRepo) GetAll() (creditCard []model.CreditCard, err error) {
	tx := r.db.Find(&creditCard)
	err = tx.Error
	return
}
