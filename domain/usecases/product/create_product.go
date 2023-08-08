package product

import (
	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/requests"
)

type CreateProductUsecase interface {
	Execute(request requests.ProductRequest) (model.Product, error)
}
