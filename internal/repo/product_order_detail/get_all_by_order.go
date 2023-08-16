package product_order_detail

import "github.com/phincon-backend/laza/domain/model"

func (r *ProductOrderDetailRepo) GetAllByOrder(orderId string) (model []model.ProductOrderDetail, err error) {
	tx := r.db.Where("order_id = ?", orderId).Find(&model)
	err = tx.Error
	return
}
