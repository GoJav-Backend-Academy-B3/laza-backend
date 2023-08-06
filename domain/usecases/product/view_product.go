package product

import "github.com/phincon-backend/laza/domain/model"

type ViewProductUsecase interface {
	Execute(offset, limit uint64) ([]model.Product, error)
}
