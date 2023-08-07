package user

import (

	"github.com/phincon-backend/laza/domain/request"

	"github.com/phincon-backend/laza/domain/requests"

	"github.com/phincon-backend/laza/helper"
)

type UpdateUserUsecase interface {

	Execute(id uint64, user request.User) *helper.Response

	Execute(id uint64, user requests.User) *helper.Response

}
