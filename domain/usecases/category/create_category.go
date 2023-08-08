package category

import (
	"github.com/phincon-backend/laza/domain/request"
	"github.com/phincon-backend/laza/domain/response"
)

type CreateCategoryUsecase interface {
	Execute(categoryRequest request.CategoryRequest) (result response.CategorySimpleResponse, err error)
}
