package auth

import (
	"github.com/phincon-backend/laza/domain/requests"
	"github.com/phincon-backend/laza/helper"
)

type ResetPasswordUserUsecase interface {
	Execute(email, code string, user requests.ResetPassword) *helper.Response
}
