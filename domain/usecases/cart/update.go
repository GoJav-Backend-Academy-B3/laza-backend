package cart

import "github.com/phincon-backend/laza/helper"

type UpdateCartUsecase interface {
	Execute(userId, productId uint64) *helper.Response
}
