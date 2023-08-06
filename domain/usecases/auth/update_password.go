package auth

import (
	"github.com/phincon-backend/laza/domain/request"
	"github.com/phincon-backend/laza/helper"
)

type UpdatePasswordUserUsecase interface {
	Execute(email, code string, user request.UpdatePassword) *helper.Response
}
