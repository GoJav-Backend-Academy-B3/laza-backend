package user

import (
	"github.com/phincon-backend/laza/domain/request"
	"github.com/phincon-backend/laza/helper"
)

type UpdateUserUsecase interface {
	Execute(id uint64, user request.User) *helper.Response
}
