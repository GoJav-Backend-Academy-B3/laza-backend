package category

import (
	"github.com/phincon-backend/laza/domain/response"
)

type GetCategoryByIdUsecase interface {
	Execute(categoryId uint64) (category response.CategorySimpleResponse, err error)
}
