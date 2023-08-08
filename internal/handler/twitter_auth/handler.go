package twitterauth

import (
	"net/http"

	handler "github.com/phincon-backend/laza/domain/handlers"
	uc "github.com/phincon-backend/laza/domain/usecases/twitter_auth"
)

type twitterAuthHandler struct {
	path           string
	pathCallback   string
	useCaseTwitter uc.TwitterAuthUsecase
}

func (h *twitterAuthHandler) GetHandlers() (hs []handler.HandlerStruct) {
	hs = append(hs,
		handler.HandlerStruct{
			Method:      http.MethodGet,
			Path:        h.path,
			HandlerFunc: h.loginTwitter,
		},
		handler.HandlerStruct{
			Method:      http.MethodGet,
			Path:        h.pathCallback,
			HandlerFunc: h.twitterCallBack,
		},
	)
	return hs
}

func NewtwitterAuthHandler(
	path string,
	pathCallback string,
	useCaseTwitter uc.TwitterAuthUsecase,

) handler.HandlerInterface {
	return &twitterAuthHandler{
		path:           path,
		pathCallback:   pathCallback,
		useCaseTwitter: useCaseTwitter,
	}
}
