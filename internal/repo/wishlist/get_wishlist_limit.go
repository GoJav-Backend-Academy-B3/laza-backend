package wishlist

import (
	"github.com/phincon-backend/laza/domain/response"
)

func (r *WishListRepo) GetCartWithLimit(userId, offset, limit uint64) (rs []response.WishProduct, err error) {
	err = r.db.Table("cart c").
		Select("p.id,p.name, p.image_url,p.price, b.name as brand_name").
		Joins("left join product p ON c.product_id = p.id").
		Joins("left join brand b on p.brand_id = b.id").
		Where("c.user_id = ?", userId).
		Limit(int(limit)).
		Offset(int(offset)).
		Find(&rs).Error
	return
}
