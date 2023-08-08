package creditcard

import "github.com/phincon-backend/laza/domain/model"

func (r *CreditCardRepo) Insert(cc model.CreditCard) (model.CreditCard, error) {
	tx := r.db.Create(&cc)
	err := tx.Error

	return cc, err
}
