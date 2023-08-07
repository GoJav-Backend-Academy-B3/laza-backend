package facebook_auth

import "github.com/phincon-backend/laza/domain/response"

type FacebookAuthUsecase interface {
	Execute(response response.FBAuthResponse) (accessToken string, err error)
}
