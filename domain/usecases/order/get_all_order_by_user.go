package order

import "github.com/phincon-backend/laza/domain/model"

type GetAllOrderByUserUsecase interface {
	Execute(orderId uint64) (orders []model.Order, err error)
}
