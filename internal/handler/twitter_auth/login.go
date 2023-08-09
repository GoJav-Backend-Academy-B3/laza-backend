package twitterauth

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/twitter"
	"github.com/phincon-backend/laza/domain/response"
	"github.com/phincon-backend/laza/helper"
)

type ProviderIndex struct {
	Providers    []string
	ProvidersMap map[string]string
}

func (h *twitterAuthHandler) loginTwitter(c *gin.Context) {
	// gothic.Store = helper.GetStore()
	gothic.Store = helper.GetStore()
	goth.UseProviders(
		twitter.New(os.Getenv("TWITTER_KEY"), os.Getenv("TWITTER_SECRET"), os.Getenv("TWITTER_REDIRECT_URI")),
	)

	if gothUser, err := gothic.CompleteUserAuth(c.Writer, c.Request); err == nil {
		rb := response.FillFromTwitter(gothUser.Email, gothUser.Name, gothUser.NickName, gothUser.RawData["profile_image_url_https"].(string))
		h.useCaseTwitter.Execute(rb).Send(c)
	} else {
		gothic.BeginAuthHandler(c.Writer, c.Request)
	}
}
