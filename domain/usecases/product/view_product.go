package product

import "github.com/phincon-backend/laza/domain/entities"

type ViewProductUsecase interface {
	Execute(offset, limit uint32) ([]entities.Product, error)
}
