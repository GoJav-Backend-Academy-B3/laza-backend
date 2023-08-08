package category

import (
	action "github.com/phincon-backend/laza/domain/repositories/category"
	"github.com/phincon-backend/laza/domain/response"
	uc "github.com/phincon-backend/laza/domain/usecases/category"
)

type getCategoryByIdUsecaseImpl struct {
	getCategoryByIdAction action.FindByIdAction
}

func (g getCategoryByIdUsecaseImpl) Execute(categoryId uint64) (category response.CategorySimpleResponse, err error) {
	result, err := g.getCategoryByIdAction.FindById(categoryId)
	if err != nil {
		return
	}
	category.Id = result.Id
	category.Category = result.Category
	return
}

func NewGetCategoryByIdUsecaseImpl(findByIdAction action.FindByIdAction) uc.GetCategoryByIdUsecase {
	return &getCategoryByIdUsecaseImpl{
		getCategoryByIdAction: findByIdAction,
	}
}
