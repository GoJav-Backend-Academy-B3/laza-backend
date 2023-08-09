package user

import (
	"github.com/phincon-backend/laza/domain/requests"
	"github.com/phincon-backend/laza/helper"
)

type UpdateUserUsecase interface {
	Execute(id uint64, user requests.UpdateUser) *helper.Response
}
