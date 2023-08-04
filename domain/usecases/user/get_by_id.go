package user

import "github.com/phincon-backend/laza/helper"

type GetByIdUserUsecase interface {
	Excute(id uint64) *helper.Response
}
