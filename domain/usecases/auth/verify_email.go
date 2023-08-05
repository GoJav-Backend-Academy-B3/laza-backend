package auth

import "github.com/phincon-backend/laza/helper"

type VerifyEmailUserUsecase interface {
	Execute(email, token string) *helper.Response
}
