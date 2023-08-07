package product

import (
	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/request"
)

type UpdateProductUsecase interface {
	Execute(id uint64, request request.ProductRequest) (model.Product, error)
}
