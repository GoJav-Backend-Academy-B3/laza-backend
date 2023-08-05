package auth

import (
	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/helper"
)

type LoginUserUsecase interface {
	Execute(user model.User) *helper.Response
}
