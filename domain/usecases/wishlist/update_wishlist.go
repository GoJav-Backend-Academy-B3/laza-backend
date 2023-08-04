package wishlist

import "github.com/phincon-backend/laza/domain/model"

type UpdateWishListUsecase interface {
	Execute(ws model.Wishlist) (data *model.Wishlist, err error)
}
