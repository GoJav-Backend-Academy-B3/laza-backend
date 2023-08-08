package category

import (
	"errors"
	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/repositories/category"
	"github.com/phincon-backend/laza/domain/request"
	"github.com/phincon-backend/laza/domain/response"
	uc "github.com/phincon-backend/laza/domain/usecases/category"
	"github.com/phincon-backend/laza/mapper"
)

type updateCategoryNameByIdUsecaseImpl struct {
	updateByIdAction category.UpdateByIdAction
}

func (ucn updateCategoryNameByIdUsecaseImpl) Execute(categoryDTO request.CategoryRequest) (category response.CategorySimpleResponse, err error) {
	var categoryModel = new(model.Category)
	categoryModel.Id = categoryDTO.Id
	categoryModel.Category = categoryDTO.Category

	if categoryDTO.Id == 0 {
		return category, errors.New("id can't be null")
	}

	updatedModel, err := ucn.updateByIdAction.Update(categoryDTO.Id, *categoryModel)
	if err != nil {
		return
	}
	category = mapper.CategoryModelToSimpleResponse(updatedModel)
	return
}

func NewUpdateCategoryNameByIdUsecaseImpl(action category.UpdateByIdAction) uc.UpdateCategoryNameByIdUsecase {
	return &updateCategoryNameByIdUsecaseImpl{
		updateByIdAction: action,
	}
}
