package auth

import "github.com/phincon-backend/laza/helper"

type ForgetPasswordUserUsecase interface {
	Execute(email string) *helper.Response
}
