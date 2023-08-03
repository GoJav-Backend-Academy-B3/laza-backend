package product

import (
	m "github.com/phincon-backend/laza/domain/model"
	d "github.com/phincon-backend/laza/domain/repositories"
	p "github.com/phincon-backend/laza/domain/usecases/product"
)

type ViewProductUsecaseImpl struct {
	getWithLimitAction d.GetWithLimitAction[m.Product]
}

// Execute implements product.ViewProductUsecase.
func (u *ViewProductUsecaseImpl) Execute(offset, limit uint64) (products []m.Product, err error) {
	products, err = u.getWithLimitAction.GetWithLimit(offset, limit)
	//TODO: Handle some error that repository might return
	return
}

func NewViewProductUsecaseImpl(
	getWithLimitAction d.GetWithLimitAction[m.Product]) p.ViewProductUsecase {
	return &ViewProductUsecaseImpl{getWithLimitAction}
}
