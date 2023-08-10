package order

import (
	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/repositories/order"
)

type GetAllOrderByUserUsecaseImpl struct {
	getAllOrderByUser order.GetAllByUser
}

func NewGetAllOrderByUserUsecaseImpl(getAllOrderByUser order.GetAllByUser) *GetAllOrderByUserUsecaseImpl {
	return &GetAllOrderByUserUsecaseImpl{getAllOrderByUser: getAllOrderByUser}
}

func (g GetAllOrderByUserUsecaseImpl) Execute(userId uint64) (orders []model.Order, err error) {
	orders, err = g.getAllOrderByUser.GetAllByUser(userId)
	return
}
