package cart

import "github.com/phincon-backend/laza/helper"

type GetCartByIdUsecase interface {
	Execute(userId uint64) *helper.Response
}
