package wishlist

import (
	"github.com/phincon-backend/laza/domain/model"
)

func (r *WishListRepo) GetById(userId any) (ws []model.Wishlist, err error) {
	err = r.db.Where("user_id =?", userId).Find(&ws).Error
	return
}
