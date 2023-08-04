package user

import "github.com/phincon-backend/laza/helper"

type DeleteUserUsecase interface {
	Excute(id uint64) *helper.Response
}
