package review

import "github.com/phincon-backend/laza/domain/model"

type GetReviewByProduct interface {
	GetProductById(productID uint64) ([]model.ProductReview, error)
	GetReviewStatsByProduct(productID uint64) (float64, int, error)
}
