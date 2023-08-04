package wishlist

import (
	"github.com/phincon-backend/laza/domain/model"
)

func (r *WishListRepo) Update(stamp string, ws *model.Wishlist) (*model.Wishlist, error) {

	if err := r.db.First(ws).Delete(ws).Error; err == nil {
		return ws, nil
	}
	ws.IsLiked = true
	err := r.db.Create(ws).Error
	if err != nil {
		return nil, err
	}
	return ws, nil
}
