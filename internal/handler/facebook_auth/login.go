package facebook_auth

import (
	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/config"
	"github.com/phincon-backend/laza/helper"
	"net/http"
)

func (fb *facebookAuthHandler) login(c *gin.Context) {
	if c.Request.Method != http.MethodGet {
		helper.GetResponse("method not allowed", http.StatusMethodNotAllowed, true).Send(c)
		return
	}

	// Create oauthState cookie
	oauthState := helper.GenerateStateOauthCookie(c)

	/*
		AuthCodeURL receive state that is a token to protect the user
		from CSRF attacks. You must always provide a non-empty string
		and validate that it matches the state query parameter
		on your redirect callback.
	*/
	redirectURL := config.FBConfig.LoginConfig.AuthCodeURL(oauthState)
	c.Redirect(http.StatusTemporaryRedirect, redirectURL)
}
