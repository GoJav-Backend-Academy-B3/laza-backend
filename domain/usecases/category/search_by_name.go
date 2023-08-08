package category

import "github.com/phincon-backend/laza/domain/response"

type SearchCategoryByNameUsecase interface {
	Execute(categoryName string) (categories []response.CategorySimpleResponse, err error)
}
