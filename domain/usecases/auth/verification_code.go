package auth

import (
	"github.com/phincon-backend/laza/domain/requests"
	"github.com/phincon-backend/laza/helper"
)

type VerificationCodeUserUsecase interface {
	Execute(user requests.VerificationCode) *helper.Response
}
