package wishlist

import "github.com/phincon-backend/laza/domain/model"

type GetWishListLimitUsecase interface {
	Execute(userId, offset, limit uint64) (result_ []model.Product, err error)
}
