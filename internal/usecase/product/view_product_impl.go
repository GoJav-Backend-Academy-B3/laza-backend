package product

import (
	e "github.com/phincon-backend/laza/domain/entities"
	d "github.com/phincon-backend/laza/domain/repositories"
	p "github.com/phincon-backend/laza/domain/usecases/product"
)

type ViewProductUsecaseImpl struct {
	getWithLimitAction d.GetWithLimitAction[e.Product]
}

// Execute implements product.ViewProductUsecase.
func (u *ViewProductUsecaseImpl) Execute(offset, limit uint64) (products []e.Product, err error) {
	products, err = u.getWithLimitAction.GetWithLimit(offset, limit)
	//TODO: Handle some error that repository might return
	return
}

func NewViewProductUsecaseImpl(
	getWithLimitAction d.GetWithLimitAction[e.Product]) p.ViewProductUsecase {
	return &ViewProductUsecaseImpl{getWithLimitAction}
}
