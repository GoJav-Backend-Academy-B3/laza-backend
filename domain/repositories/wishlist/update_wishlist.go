package wishlist

import "github.com/phincon-backend/laza/domain/model"

type UpdateWishListAction interface {
	UpdateWishList(model model.Wishlist) (value any, err error)
}
