package transaction

import "github.com/phincon-backend/laza/domain/model"

func (r *TransactionRepo) Delete(id any) (err error) {

	tx := r.db.Delete(model.Transaction{}, id.(string))
	err = tx.Error

	return err
}
