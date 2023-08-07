package user

import "github.com/phincon-backend/laza/helper"

type GetByIdUserUsecase interface {
	Execute(id uint64) *helper.Response
}
