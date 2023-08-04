package facebook_auth

import "github.com/phincon-backend/laza/domain/usecases/facebook_auth"

type facebookAuthUsecaseImpl struct {
}

func (fb *facebookAuthUsecaseImpl) Execute() (err error) {
	return
}

func NewFacebookAuthUsecase() facebook_auth.FacebookAuthUsecase {
	return &facebookAuthUsecaseImpl{}
}
