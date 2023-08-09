package auth

import "github.com/phincon-backend/laza/helper"

type ForgotPasswordUserUsecase interface {
	Execute(email string) *helper.Response
}
