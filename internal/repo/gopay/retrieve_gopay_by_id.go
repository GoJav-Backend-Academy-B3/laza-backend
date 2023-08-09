package gopay

import "github.com/phincon-backend/laza/domain/model"

func (r *GopayRepo) GetById(id any) (gopay model.Gopay, err error) {
	idConv := id.(uint64)

	tx := r.db.First(&gopay, idConv)
	err = tx.Error
	return
}