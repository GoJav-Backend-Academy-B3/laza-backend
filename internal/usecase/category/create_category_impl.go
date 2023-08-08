package category

import (
	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/repositories"
	"github.com/phincon-backend/laza/domain/request"
	"github.com/phincon-backend/laza/domain/response"
	usecase "github.com/phincon-backend/laza/domain/usecases/category"
	"github.com/phincon-backend/laza/internal/repo/category"
)

type createCategoryUsecaseImpl struct {
	insertCategoryAction repositories.InsertAction[model.Category]
}

func (cc *createCategoryUsecaseImpl) Execute(categoryRequest request.CategoryRequest) (result response.CategorySimpleResponse, err error) {
	var category model.Category
	category.SetCategory(categoryRequest.GetCategory())

	category, err = cc.insertCategoryAction.Insert(category)
	return
}

func NewCreateCategoryUsecaseImpl(categoryRepo category.CategoryRepo) usecase.CreateCategoryUsecase {
	return &createCategoryUsecaseImpl{
		insertCategoryAction: &categoryRepo,
	}
}
