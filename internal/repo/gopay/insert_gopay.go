package gopay

import "github.com/phincon-backend/laza/domain/model"

func (r *GopayRepo) Insert(gp model.Gopay) (model.Gopay, error) {
	tx := r.db.Create(&gp)
	err := tx.Error

	return gp, err
}