package user

import "github.com/phincon-backend/laza/helper"

type DeleteUserUsecase interface {
	Execute(id uint64) *helper.Response
}
