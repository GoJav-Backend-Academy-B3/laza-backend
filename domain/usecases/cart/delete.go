package cart

import "github.com/phincon-backend/laza/helper"

type DeleteCartUsecase interface {
	Execute(userId uint64, productId uint64) *helper.Response
}
