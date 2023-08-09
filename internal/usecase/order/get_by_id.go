package order

import (
	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/repositories"
)

type GetByIdUsecase struct {
	getByIdAction repositories.GetByIdAction[model.Order]
}

func NewGetByIdUsecase(getByIdAction repositories.GetByIdAction[model.Order]) *GetByIdUsecase {
	return &GetByIdUsecase{getByIdAction: getByIdAction}
}

// Execute implements product.SearchProductByNameUsecase.
func (u *GetByIdUsecase) Execute(orderId any) (order model.Order, err error) {
	order, err = u.getByIdAction.GetById(orderId)
	return
}
