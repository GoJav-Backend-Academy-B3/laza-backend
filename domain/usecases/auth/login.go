package auth

import (
	"github.com/phincon-backend/laza/domain/requests"
	"github.com/phincon-backend/laza/helper"
)

type LoginUserUsecase interface {
	Execute(user requests.Login) *helper.Response
}
