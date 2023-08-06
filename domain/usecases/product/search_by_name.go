package product

import "github.com/phincon-backend/laza/domain/model"

type SearchProductByNameUsecase interface {
	Execute(keyword string, offset, limit uint64) ([]model.Product, error)
}
