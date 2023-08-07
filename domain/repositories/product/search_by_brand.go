package product

import "github.com/phincon-backend/laza/domain/model"

type SearchByBrandAction interface {
	SearchByBrand(brand string, offset, limit uint64) ([]model.Product, error)
}
