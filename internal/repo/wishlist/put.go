package wishlist

import (
	"github.com/phincon-backend/laza/domain/model"
)

func (r *WishListRepo) Update(stamp string, ws model.Wishlist) (rs model.Wishlist, err error) {

	if err = r.db.First(&rs, "user_id = ? AND product_id =?", ws.UserId, ws.ProductId).Error; rs != (model.Wishlist{}) {
		if err != nil {
			return
		}
		valueIsLiked := false
		switch {
		case rs.IsLiked == true:
			valueIsLiked = false
		case ws.IsLiked == false:
			valueIsLiked = true
		}
		// Single Columns
		tx := r.db.Model(&ws).Where("user_id = ? AND product_id = ?", ws.UserId, ws.ProductId).Update("is_liked", valueIsLiked).Scan(&rs)
		err = tx.Error
	} else {
		ws.IsLiked = true
		err = r.db.Create(&ws).Scan(&rs).Error
	}
	return
}
