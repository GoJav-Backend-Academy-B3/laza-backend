package review

import "github.com/phincon-backend/laza/domain/model"

type GetReviewByProduct interface {
	GetProductById(productID uint64) ([]model.ProductReview, error)
	GetReviewStatsByProduct(productID uint64) (float64, int, error)
}
type GetWithLimitReviewsAction interface {
	GetWithLimit(offset, limit uint64, productID uint64) ([]model.ProductReview, error)
}

// type GetReviewWithLimitAction interface {
// 	GetWithLimit(offset, limit uint64, productID uint64) ([]model.ProductReview, error)
// }
