package category

import (
	"github.com/phincon-backend/laza/domain/request"
	"github.com/phincon-backend/laza/domain/response"
)

type UpdateCategoryNameByIdUsecase interface {
	Execute(categoryDTO request.CategoryRequest) (category response.CategorySimpleResponse, err error)
}
