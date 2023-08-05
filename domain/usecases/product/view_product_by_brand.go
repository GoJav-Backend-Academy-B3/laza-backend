package product

import "github.com/phincon-backend/laza/domain/model"

type ViewProductByBrandUsecase interface {
	Execute(brand string, offset, limit uint64) ([]model.Product, error)
}
