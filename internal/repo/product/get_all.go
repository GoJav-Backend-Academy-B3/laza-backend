package product

import "github.com/phincon-backend/laza/domain/model"

func (r *ProductRepo) GetAll() (es []model.Product, err error) {
	tx := r.db.Find(&es)
	err = tx.Error
	return
}
