package facebook_auth

import (
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
	"github.com/phincon-backend/laza/helper"
	"log"
	"net/http"
	"sort"
)

type ProviderIndex struct {
	Providers    []string
	ProvidersMap map[string]string
}

func (fb *facebookAuthHandler) loginTwitter(c *gin.Context) {
	gothic.Store = helper.GetStore()
	m := make(map[string]string)
	m["twitter"] = "Twitter"

	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	if gothUser, err := gothic.CompleteUserAuth(c.Writer, c.Request); err == nil {
		log.Println(gothUser)
		helper.GetResponse(gothUser, http.StatusOK, false).Send(c)
	} else {
		gothic.BeginAuthHandler(c.Writer, c.Request)
	}
}
