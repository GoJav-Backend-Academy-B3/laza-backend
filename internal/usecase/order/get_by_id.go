package order

import (
	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/repositories"
)

type GetByIdUsecase struct {
	getOrder repositories.GetByIdAction[model.Order]
}

func NewGetByIdUsecase(getOrder repositories.GetByIdAction[model.Order]) *GetByIdUsecase {
	return &GetByIdUsecase{getOrder: getOrder}
}

// Execute implements product.SearchProductByNameUsecase.
func (u *GetByIdUsecase) Execute(orderId string) (order model.Order, err error) {
	order, err = u.getOrder.GetById(orderId)
	return
}
