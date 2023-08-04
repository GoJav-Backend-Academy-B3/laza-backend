package wishlist

import (
	m "github.com/phincon-backend/laza/domain/model"
	d "github.com/phincon-backend/laza/domain/repositories"
	p "github.com/phincon-backend/laza/domain/usecases/wishlist"
)

type UpdateWishListUsecaseImpl struct {
	updateWishlist d.UpdateAction[m.Wishlist]
}

func (u *UpdateWishListUsecaseImpl) Execute(ws m.Wishlist) (wishlist m.Wishlist, err error) {
	wishlist, err = u.updateWishlist.Update("", ws)

	return
}

func NewUpdateWishListUsecaseImpl(uw d.UpdateAction[m.Wishlist]) p.UpdateWishListUsecase {
	return &UpdateWishListUsecaseImpl{
		updateWishlist: uw,
	}
}
