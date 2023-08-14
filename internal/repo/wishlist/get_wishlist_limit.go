package wishlist

import (
	"github.com/phincon-backend/laza/domain/model"
)

func (r *WishListRepo) GetWishlistProductLimit(userId, offset, limit uint64) (rs []model.Product, err error) {

	tx := r.db.
		Table("product p").
		Select("p.id, p.name, p.image_url,p.price,p.created_at").
		Joins("JOIN wishlist w on w.product_id = p.id").
		Where("w.user_id=? AND p.deleted_at is null", userId).
		Offset(int(offset)).
		Limit(int(limit)).
		Find(&rs)

	err = tx.Error
	return
}
