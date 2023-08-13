package wishlist

import (
	"github.com/phincon-backend/laza/domain/model"
	"gorm.io/gorm"
)

func (r *WishListRepo) UpdateWishList(user_id any, product_id any) (value any, err error) {
	userId := user_id.(uint64)
	productId := product_id.(uint64)
	var ws model.Wishlist

	tx := r.db.Where("user_id = ? AND product_id = ?", userId, productId).First(&ws)
	err = tx.Error
	value = ""

	if err == gorm.ErrRecordNotFound {
		tx = r.db.Create(&model.Wishlist{
			UserId:    userId,
			ProductId: productId},
		)
		err = tx.Error
		value = "successfully added wishlist"

	} else {
		tx = r.db.Delete(&model.Wishlist{
			UserId:    userId,
			ProductId: productId},
		)
		value = "successfully delete wishlist"
		err = tx.Error
	}

	return
}
