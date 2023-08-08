package product

import (
	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/requests"
)

type UpdateProductUsecase interface {
	Execute(id uint64, request requests.ProductRequest) (model.Product, error)
}
