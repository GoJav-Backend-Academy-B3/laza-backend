package category

import "github.com/phincon-backend/laza/domain/response"

type ViewCategoryUsecase interface {
	Execute() (categories []response.CategorySimpleResponse, err error)
}
