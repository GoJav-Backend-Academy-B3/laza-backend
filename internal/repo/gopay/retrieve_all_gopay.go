package gopay

import "github.com/phincon-backend/laza/domain/model"

func (r *GopayRepo) GetAll() (gopay []model.Gopay, err error) {
	tx := r.db.Find(&gopay)
	err = tx.Error
	return
}
