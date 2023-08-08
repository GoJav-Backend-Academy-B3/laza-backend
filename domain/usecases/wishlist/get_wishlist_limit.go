package wishlist

import "github.com/phincon-backend/laza/helper"

type GetWishListLimitUsecase interface {
	Execute(userId, offset, limit uint64) *helper.Response
}
