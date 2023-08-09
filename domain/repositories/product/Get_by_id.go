package product

import "github.com/phincon-backend/laza/domain/model"

type GetProductByIdAction interface {
	GetProductById(productID uint64) ([]model.ProductReview, error)
}
