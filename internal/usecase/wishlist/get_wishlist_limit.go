package wishlist

import (
	"github.com/phincon-backend/laza/domain/model"

	d "github.com/phincon-backend/laza/domain/repositories/wishlist"
	p "github.com/phincon-backend/laza/domain/usecases/wishlist"
)

type getWishlistLimitUsecase struct {
	getWishlistLimitRepo d.GetWishlistProductAction
}

func (uc *getWishlistLimitUsecase) Execute(userId, offset, limit uint64) (result_ []model.Product, err error) {

	result_, err = uc.getWishlistLimitRepo.GetWishlistProductLimit(userId, offset, limit)

	return

}

func NewgetWishlistLimitUsecase(gwr d.GetWishlistProductAction) p.GetWishListLimitUsecase {
	return &getWishlistLimitUsecase{
		getWishlistLimitRepo: gwr,
	}
}
