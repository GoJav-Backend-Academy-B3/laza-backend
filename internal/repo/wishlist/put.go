package wishlist

import (
	"github.com/phincon-backend/laza/domain/model"
	"gorm.io/gorm"
)

func (r *WishListRepo) Update(stamp any, ws model.Wishlist) (rs model.Wishlist, err error) {

	if err = r.db.Where("user_id = ? AND product_id = ?", ws.UserId, ws.ProductId).First(&ws).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			err = r.db.Create(&ws).Error
			rs = ws
			if err != nil {
				return
			}
		}
		return
	}
	err = r.db.Delete(&ws).Error
	if err != nil {
		return
	}
	return
}
