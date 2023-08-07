package facebook_auth

import (
	handler "github.com/phincon-backend/laza/domain/handlers"
	"github.com/phincon-backend/laza/domain/usecases/facebook_auth"
	"net/http"
)

type facebookAuthHandler struct {
	loginPath           string
	callbackPath        string
	facebookAuthUsecase facebook_auth.FacebookAuthUsecase
}

func (fb *facebookAuthHandler) GetHandlers() (hs []handler.HandlerStruct) {
	hs = append(hs, handler.HandlerStruct{
		Method:      http.MethodGet,
		Path:        fb.loginPath,
		HandlerFunc: fb.loginTwitter,
	}, handler.HandlerStruct{
		Method:      http.MethodGet,
		Path:        fb.callbackPath,
		HandlerFunc: fb.TwitterCallback,
	})
	return
}

func NewFacebookAuthHandler(
	loginPath string,
	callbackPath string,
	facebookAuthUsecase facebook_auth.FacebookAuthUsecase) handler.HandlerInterface {
	return &facebookAuthHandler{
		loginPath:           loginPath,
		callbackPath:        callbackPath,
		facebookAuthUsecase: facebookAuthUsecase,
	}
}
