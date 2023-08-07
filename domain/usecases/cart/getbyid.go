package cart

import "github.com/phincon-backend/laza/helper"

type GetCartByIdUsecase interface {
	Execute(userd any) *helper.Response
}
