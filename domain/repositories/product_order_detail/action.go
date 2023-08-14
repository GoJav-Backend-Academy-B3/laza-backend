package product_order_detail

import "github.com/phincon-backend/laza/domain/model"

type GetByOrder interface {
	GetAllByOrder(orderId string) (model []model.ProductOrderDetail, err error)
}
