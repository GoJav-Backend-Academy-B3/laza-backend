package order

import (
	"github.com/phincon-backend/laza/domain/model"
)

type GetOrderByIdUsecase interface {
	Execute(orderId string) (order model.Order, err error)
}
