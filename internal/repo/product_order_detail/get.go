package product_order_detail

import "github.com/phincon-backend/laza/domain/model"

func (r *ProductOrderDetailRepo) GetById(id any) (order model.ProductOrderDetail, err error) {
	tx := r.db.First(&order, "id = ?", id)
	err = tx.Error
	return
}
