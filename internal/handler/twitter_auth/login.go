package twitterauth

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
	"github.com/phincon-backend/laza/helper"
)

type ProviderIndex struct {
	Providers    []string
	ProvidersMap map[string]string
}

func (h *twitterAuthHandler) loginTwitter(c *gin.Context) {
	gothic.Store = helper.GetStore()

	if gothUser, err := gothic.CompleteUserAuth(c.Writer, c.Request); err == nil {
		log.Println(gothUser)
		helper.GetResponse(gothUser, http.StatusOK, false).Send(c)
	} else {
		gothic.BeginAuthHandler(c.Writer, c.Request)
	}
}
