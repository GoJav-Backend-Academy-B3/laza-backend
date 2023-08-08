package category

import (
	"github.com/phincon-backend/laza/domain/repositories/category"
	"github.com/phincon-backend/laza/domain/response"
	uc "github.com/phincon-backend/laza/domain/usecases/category"
	"github.com/phincon-backend/laza/mapper"
)

type viewCategoryUsecaseImpl struct {
	findAllAction category.FindAllAction
}

func (v viewCategoryUsecaseImpl) Execute() (categories []response.CategorySimpleResponse, err error) {
	records, err := v.findAllAction.FindAll()
	if err != nil {
		return
	}
	for _, item := range records {
		categories = append(categories, mapper.CategoryModelToSimpleResponse(item))
	}
	return
}

func NewViewCategoryUsecaseImpl(action category.FindAllAction) uc.ViewCategoryUsecase {
	return &viewCategoryUsecaseImpl{
		findAllAction: action,
	}
}
