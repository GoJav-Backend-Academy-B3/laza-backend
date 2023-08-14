package wishlist

import "github.com/phincon-backend/laza/domain/requests"

type UpdateWishListUsecase interface {
	Execute(userId uint64, rb requests.WishlistRequest) (value any, err error)
}
