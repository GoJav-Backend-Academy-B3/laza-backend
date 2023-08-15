package wishlist

import (
	"github.com/phincon-backend/laza/domain/model"
)

type GetWishlistProductAction interface {
	GetWishlistProductLimit(userId, offset, limit uint64) (rs []model.Product, err error)
}
