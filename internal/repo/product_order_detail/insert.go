package product_order_detail

import "github.com/phincon-backend/laza/domain/model"

func (r *ProductOrderDetailRepo) Insert(model model.ProductOrderDetail) (model.ProductOrderDetail, error) {
	tx := r.db.Create(&model)
	err := tx.Error

	return model, err
}
