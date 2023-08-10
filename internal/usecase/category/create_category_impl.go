package category

import (
	"github.com/go-playground/validator/v10"
	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/repositories"
	"github.com/phincon-backend/laza/domain/requests"
	"github.com/phincon-backend/laza/domain/response"
	usecase "github.com/phincon-backend/laza/domain/usecases/category"
	"github.com/phincon-backend/laza/internal/repo/category"
	"github.com/phincon-backend/laza/mapper"
)

type createCategoryUsecaseImpl struct {
	insertCategoryAction repositories.InsertAction[model.Category]
}

func (cc *createCategoryUsecaseImpl) Execute(categoryRequest requests.CategoryRequest) (result response.CategorySimpleResponse, err error) {
	var categoryModel = new(model.Category)
	err = validator.New().Struct(categoryRequest)
	if err != nil {
		return
	}
	categoryModel.SetCategory(categoryRequest.GetCategory())

	res, err := cc.insertCategoryAction.Insert(*categoryModel)
	result = mapper.CategoryModelToSimpleResponse(res)
	return
}

func NewCreateCategoryUsecaseImpl(categoryRepo category.CategoryRepo) usecase.CreateCategoryUsecase {
	return &createCategoryUsecaseImpl{
		insertCategoryAction: &categoryRepo,
	}
}
