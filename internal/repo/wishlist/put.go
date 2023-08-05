package wishlist

import (
	"github.com/phincon-backend/laza/domain/model"
)

func (r *WishListRepo) Update(stamp any, ws *model.Wishlist) (*model.Wishlist, error) {

	r.db.Where("user_id = ? AND product_id = ?", ws.UserId, ws.ProductId).Find(ws)
	if ws.IsLiked == false {
		ws.IsLiked = true
		err := r.db.Create(ws).Error
		if err != nil {
			return nil, err
		}
		return ws, nil
	}
	err := r.db.Delete(ws).Error
	if err != nil {
		return nil, err
	}
	return ws, nil
}
