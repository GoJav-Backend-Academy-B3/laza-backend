package product

import "github.com/phincon-backend/laza/domain/model"

func (r *ProductRepo) Insert(e model.Product) (model.Product, error) {
	tx := r.db.Create(&e)
	return e, tx.Error
}
