package wishlist

import "github.com/phincon-backend/laza/domain/response"

type GetCartWithLimitAction interface {
	GetCartWithLimit(userId, offset, limit uint64) (rs []response.CartPorduct, err error)
}
