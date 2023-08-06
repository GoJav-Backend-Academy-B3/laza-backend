package product

import "github.com/phincon-backend/laza/domain/model"

type SearchByNameAction interface {
	SearchByName(keyword string, offset, limit uint64) ([]model.Product, error)
}
