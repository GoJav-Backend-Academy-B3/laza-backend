package wishlist

import (
	"github.com/phincon-backend/laza/domain/response"
)

type GetWishListUsecase interface {
	Execute(userId uint64) (data *[]response.WishlistProduct, err error)
}
