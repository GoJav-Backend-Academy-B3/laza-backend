package wishlist

import (
	wir "github.com/phincon-backend/laza/domain/repositories/wishlist"
	p "github.com/phincon-backend/laza/domain/usecases/wishlist"
)

type UpdateWishListUsecaseImpl struct {
	updateWishlist wir.UpdateWishListAction
}

func (u *UpdateWishListUsecaseImpl) Execute(userId, productId uint64) (value any, err error) {

	value, err = u.updateWishlist.UpdateWishList(userId, productId)

	return
}

func NewUpdateWishListUsecaseImpl(
	updateWishlist wir.UpdateWishListAction,
) p.UpdateWishListUsecase {
	return &UpdateWishListUsecaseImpl{
		updateWishlist: updateWishlist,
	}
}
