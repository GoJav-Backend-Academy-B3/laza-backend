package user

import (
	"github.com/phincon-backend/laza/domain/requests"
	"github.com/phincon-backend/laza/helper"
)

type ChangePasswordUserUsecase interface {
	Execute(id uint64, user requests.ChangePassword) *helper.Response
}
