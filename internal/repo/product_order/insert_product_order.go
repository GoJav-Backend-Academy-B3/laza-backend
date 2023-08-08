package product_order

import "github.com/phincon-backend/laza/domain/model"

func (r *ProductOrderRepo) Insert(po model.ProductOrder) (model.ProductOrder, error) {
	tx := r.db.Create(&po)
	err := tx.Error

	return po, err
}
