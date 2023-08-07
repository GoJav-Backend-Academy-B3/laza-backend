package auth

import "github.com/phincon-backend/laza/helper"

type LoginGoogleUserUsecase interface {
	Execute(user *helper.GoogleUserResult) *helper.Response
}
