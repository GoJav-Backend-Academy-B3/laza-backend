package creditcard

import (
	"errors"
	"github.com/phincon-backend/laza/domain/model"
)

func (r *CreditCardRepo) Update(id any, cc model.CreditCard) (model.CreditCard, error) {
	if id.(uint64) != cc.Id {
		return cc, errors.New("id param and model isn't same")
	}

	tx := r.db.Save(&cc)
	err := tx.Error

	return cc, err
}
