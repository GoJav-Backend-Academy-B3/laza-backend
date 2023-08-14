package wishlist

import (
	"github.com/phincon-backend/laza/domain/model"
	"gorm.io/gorm"
)

func (r *WishListRepo) UpdateWishList(md model.Wishlist) (value any, err error) {

	var ws model.Wishlist

	tx := r.db.Where("user_id = ? AND product_id = ?", md.UserId, md.ProductId).First(&ws)
	err = tx.Error
	value = ""

	if err == gorm.ErrRecordNotFound {
		tx = r.db.Create(&md)
		err = tx.Error
		value = "successfully added wishlist"

	} else {
		tx = r.db.Delete(&md)
		value = "successfully delete wishlist"
		err = tx.Error
	}

	return
}
