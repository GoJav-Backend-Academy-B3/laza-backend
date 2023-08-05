package auth

import "github.com/phincon-backend/laza/helper"

type ResendEmailUserUsecase interface {
	Execute(email string) *helper.Response
}
