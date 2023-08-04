package user

import (
	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/helper"
)

type InsertUserUsecase interface {
	Excute(user model.User) *helper.Response
}
