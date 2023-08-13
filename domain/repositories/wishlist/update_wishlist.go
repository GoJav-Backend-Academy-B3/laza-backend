package wishlist

type UpdateWishListAction interface {
	UpdateWishList(userId, productId any) (value any, err error)
}
