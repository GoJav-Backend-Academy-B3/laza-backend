package order

import (
	"github.com/phincon-backend/laza/domain/model"
)

type GetOrderByIdUsecase interface {
	Execute(orderId any) (order model.Order, err error)
}
