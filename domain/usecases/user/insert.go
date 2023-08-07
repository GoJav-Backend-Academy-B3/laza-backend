package user

import (
	"github.com/phincon-backend/laza/domain/requests"
	"github.com/phincon-backend/laza/helper"
)

type InsertUserUsecase interface {
	Execute(user requests.User) *helper.Response
}
