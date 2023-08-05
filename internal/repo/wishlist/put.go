package wishlist

import (
	"github.com/phincon-backend/laza/domain/model"
)

func (r *WishListRepo) Update(stamp any, ws model.Wishlist) (rs model.Wishlist, err error) {

	r.db.Where("user_id = ? AND product_id = ?", ws.UserId, ws.ProductId).Find(&ws)
	if ws.IsLiked == false {
		ws.IsLiked = true
		err = r.db.Create(&ws).Scan(&rs).Error
		if err != nil {
			return
		}
		return
	}
	rs = ws
	err = r.db.Delete(&ws).Error
	if err != nil {
		return
	}
	return
}
