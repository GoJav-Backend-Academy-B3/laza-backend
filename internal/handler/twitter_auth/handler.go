package twitterauth

import (
	"net/http"

	handler "github.com/phincon-backend/laza/domain/handlers"
)

type twitterAuthHandler struct {
	path         string
	pathCallback string
	// useCaseTwitter uc.TwitterAuthUsecase
}

func (h *twitterAuthHandler) GetHandlers() (hs []handler.HandlerStruct) {
	hs = append(hs,
		handler.HandlerStruct{
			Method:      http.MethodGet,
			Path:        h.path,
			HandlerFunc: h.TwitterCallBack,
		},
		handler.HandlerStruct{
			Method:      http.MethodGet,
			Path:        h.pathCallback,
			HandlerFunc: h.loginTwitter,
		},
	)
	return hs
}

func NewtwitterAuthHandler(path string,
	pathCallback string,
) handler.HandlerInterface {
	return &twitterAuthHandler{
		path:         path,
		pathCallback: pathCallback,
		//useCaseTwitter: useCaseTwitter,
	}
}
