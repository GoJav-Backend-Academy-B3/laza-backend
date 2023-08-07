package user

import (
	"github.com/phincon-backend/laza/domain/request"
	"github.com/phincon-backend/laza/helper"
)

type InsertUserUsecase interface {
	Execute(user request.User) *helper.Response
}
