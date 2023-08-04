package user

import "github.com/phincon-backend/laza/helper"

type GetWithLimitUserUsecase interface {
	Excute(limit, offset uint64) *helper.Response
}
