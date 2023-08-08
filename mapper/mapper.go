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

func ProductModelToProductResponse(source model.Product) response.Product {
	var result = new(response.Product)
	result.Id = source.Id
	result.Name = source.Name
	result.Description = source.Description
	result.ImageUrl = source.ImageUrl
	result.Price = source.Price
	result.CategoryId = source.CategoryId
	result.BrandId = source.BrandId
	result.CreatedAt = source.CreatedAt
	result.UpdatedAt = source.UpdatedAt
	for _, v := range source.Sizes {
		result.Size = append(result.Size, v.Size)
	}
	return *result
}
