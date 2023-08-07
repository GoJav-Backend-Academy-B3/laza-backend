package facebook_auth

import (
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
	"github.com/phincon-backend/laza/helper"
	"net/http"
)

func (fb *facebookAuthHandler) TwitterCallback(c *gin.Context) {
	//goth.UseProviders(
	//	twitter.New(os.Getenv("TWITTER_KEY"), os.Getenv("TWITTER_SECRET"), "http://127.0.0.1:8080/auth/twitter/callback"),
	//)
	gothic.Store = helper.GetStore()
	user, err := gothic.CompleteUserAuth(c.Writer, c.Request)

	if err != nil {
		panic(err)
		helper.GetResponse(err.Error(), http.StatusUnauthorized, true).Send(c)
		return
	}

	// send back response to browse
	helper.GetResponse(user, http.StatusOK, false).Send(c)
}
