package product

import (
	"github.com/phincon-backend/laza/domain/model"
	action "github.com/phincon-backend/laza/domain/repositories/product"
	usecase "github.com/phincon-backend/laza/domain/usecases/product"
)

type SearchByNameUsecaseImpl struct {
	searchByNameAction action.SearchByNameAction
}

// Execute implements product.SearchProductByNameUsecase.
func (u *SearchByNameUsecaseImpl) Execute(keyword string, offset uint64, limit uint64) (products []model.Product, err error) {
	products, err = u.searchByNameAction.SearchByName(keyword, offset, limit)
	//TODO: Handle some error that repository might return
	return
}

func NewSearchProductUsecaseImpl(
	action action.SearchByNameAction,
) usecase.SearchProductByNameUsecase {
	return &SearchByNameUsecaseImpl{
		searchByNameAction: action,
	}
}
