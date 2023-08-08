package product

import (
	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/request"
)

type CreateProductUsecase interface {
	Execute(request request.ProductRequest) (model.Product, error)
}
