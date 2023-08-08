package category

import (
	"github.com/phincon-backend/laza/domain/requests"
	"github.com/phincon-backend/laza/domain/response"
)

type UpdateCategoryNameByIdUsecase interface {
	Execute(categoryDTO requests.CategoryRequest) (category response.CategorySimpleResponse, err error)
}
