package creditcard

import (
	"errors"
	"github.com/phincon-backend/laza/domain/model"
)

func (r *CreditCardRepo) GetById(id any) (cc model.CreditCard, err error) {
	idConv, ok := id.(uint64)

	if ok == false {
		return cc, errors.New("error when convert id to uint64")
	}

	tx := r.db.First(&cc, idConv)
	err = tx.Error
	return
}
