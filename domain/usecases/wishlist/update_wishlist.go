package wishlist

type UpdateWishListUsecase interface {
	Execute(userId, productId uint64) (value any, err error)
}
