package product

import (
	"github.com/phincon-backend/laza/domain/model"
	action "github.com/phincon-backend/laza/domain/repositories/product"
	usecase "github.com/phincon-backend/laza/domain/usecases/product"
)

type ViewProductByBrandUsecaseImpl struct {
	searchByBrandAction action.SearchByBrandAction
}

func (pb *ViewProductByBrandUsecaseImpl) Execute(brand string, offset uint64, limit uint64) (products []model.Product, err error) {
	products, err = pb.searchByBrandAction.SearchByBrand(brand, offset, limit)
	return
}
func NewViewProductByBrandUsecaseImpl(action action.SearchByBrandAction) usecase.ViewProductByBrandUsecase {
	return &ViewProductByBrandUsecaseImpl{
		searchByBrandAction: action,
	}
}
