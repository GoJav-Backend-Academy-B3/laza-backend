package mapper

import (
	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/response"
)

func CategoryModelToSimpleResponse(source model.Category) response.CategorySimpleResponse {
	var result = new(response.CategorySimpleResponse)
	result.Category = source.Category
	result.Id = source.Id
	return *result
}
