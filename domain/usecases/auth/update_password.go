package auth

import (
	"github.com/phincon-backend/laza/domain/requests"
	"github.com/phincon-backend/laza/helper"
)

type UpdatePasswordUserUsecase interface {
	Execute(email, code string, user requests.UpdatePassword) *helper.Response
}
