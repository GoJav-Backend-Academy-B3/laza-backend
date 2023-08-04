package wishlist

import (
	"github.com/phincon-backend/laza/domain/model"
)

func (r *WishListRepo) GetById(userId any) (ws *[]model.Wishlist, err error) {
	if err = r.db.Where("user_id =?", userId).Find(ws).Error; err != nil {
		return nil, err
	}
	return
}
