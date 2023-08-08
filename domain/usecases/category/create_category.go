package category

import (
	"github.com/phincon-backend/laza/domain/requests"
	"github.com/phincon-backend/laza/domain/response"
)

type CreateCategoryUsecase interface {
	Execute(categoryRequest requests.CategoryRequest) (result response.CategorySimpleResponse, err error)
}
