package facebook_auth

import (
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
	"github.com/phincon-backend/laza/helper"
	"log"
	"net/http"
)

type ProviderIndex struct {
	Providers    []string
	ProvidersMap map[string]string
}

func (fb *facebookAuthHandler) loginTwitter(c *gin.Context) {
	gothic.Store = helper.GetStore()
	if gothUser, err := gothic.CompleteUserAuth(c.Writer, c.Request); err == nil {
		log.Println(gothUser)
		helper.GetResponse(gothUser, http.StatusOK, false).Send(c)
	} else {
		gothic.BeginAuthHandler(c.Writer, c.Request)
	}
}
