package cart

import "github.com/phincon-backend/laza/helper"

type GetCartOrderInfoUsecase interface {
	Execute(UserId uint64) *helper.Response
}
