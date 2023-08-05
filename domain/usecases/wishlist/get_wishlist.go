package wishlist

import "github.com/phincon-backend/laza/helper"

type GetWishListUsecase interface {
	Execute(userId uint64) *helper.Response
}
