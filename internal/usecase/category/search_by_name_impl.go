package category

import (
	"github.com/phincon-backend/laza/domain/repositories/category"
	"github.com/phincon-backend/laza/domain/response"
	usecase "github.com/phincon-backend/laza/domain/usecases/category"
	"github.com/phincon-backend/laza/mapper"
)

type searchCategoryByNameUsecaseImpl struct {
	findByNameAction category.FindByNameAction
}

func (s searchCategoryByNameUsecaseImpl) Execute(categoryName string) (categories []response.CategorySimpleResponse, err error) {
	records, err := s.findByNameAction.FindByName(categoryName)
	if err != nil {
		return nil, err
	}

	for _, item := range records {
		categories = append(categories, mapper.CategoryModelToSimpleResponse(item))
	}
	return
}

func NewSearchCategoryByNameUsecaseImpl(action category.FindByNameAction) usecase.SearchCategoryByNameUsecase {
	return &searchCategoryByNameUsecaseImpl{
		findByNameAction: action,
	}
}
