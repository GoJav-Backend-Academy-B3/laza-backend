package order

import "github.com/phincon-backend/laza/domain/model"

type GetByUser interface {
	GetById(orderId string, userId uint64) (order model.Order, err error)
}
