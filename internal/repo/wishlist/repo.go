package wishlist

import "gorm.io/gorm"

type WishListRepo struct {
	db *gorm.DB
}

func NewWishList(db *gorm.DB) *WishListRepo {
	return &WishListRepo{
		db: db,
	}
}
