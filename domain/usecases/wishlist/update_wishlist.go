package wishlist

import (
	"github.com/phincon-backend/laza/helper"
)

type UpdateWishListUsecase interface {
	Execute(userId, productId uint64) *helper.Response
}
