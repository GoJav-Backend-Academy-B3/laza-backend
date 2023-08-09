package review

import (
	"github.com/phincon-backend/laza/domain/model"
)

func (r *ReviewRepo) GetWithLimit(offset, limit uint64, productID uint64) (es []model.ProductReview, err error) {
	tx := r.db.
		Model(&model.Review{}).
		Select("review.id, review.comment, review.rating, review.created_at, users.full_name, users.image_url").
		Joins("left join users on review.user_id = users.id").
		Where("product_id = ?", productID).Offset(int(offset)).Limit(int(limit)).
		Scan(&es)
	err = tx.Error
	return
}
