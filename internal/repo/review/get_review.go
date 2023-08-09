package review

import (
	"github.com/phincon-backend/laza/domain/model"
)

func (r *ReviewRepo) GetProductById(productID uint64) (reviews []model.ProductReview, err error) {
	tx := r.db.
		Model(&model.Review{}).
		Select("review.id, review.comment, review.rating, review.created_at, users.full_name, users.image_url").
		Joins("left join users on review.user_id = users.id").
		Where("product_id = ?", productID).
		Scan(&reviews)
	err = tx.Error
	return
}
