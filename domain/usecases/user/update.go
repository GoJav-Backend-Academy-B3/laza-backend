package user

import (
	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/helper"
)

type UpdateUserUsecase interface {
	Excute(id uint64, user model.User) *helper.Response
}