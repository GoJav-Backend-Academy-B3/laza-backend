package user

import (
	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/helper"
)

type InsertUserUsecase interface {
	Execute(user model.User) *helper.Response
}
