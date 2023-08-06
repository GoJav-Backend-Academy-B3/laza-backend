package user

import "github.com/phincon-backend/laza/helper"

type GetWithLimitUserUsecase interface {
	Execute(limit, offset uint64) *helper.Response
}
