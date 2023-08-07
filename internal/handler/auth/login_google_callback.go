package auth

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/config"
	"github.com/phincon-backend/laza/helper"
)

func (h *authHandler) loginGoogleCallback(c *gin.Context) {
	// get oauth state from cookie for this user
	oauthState, _ := c.Request.Cookie("oauthstate")
	state := c.Query("state")
	code := c.Query("code")

	// ERROR : Invalid OAuth State
	if state != oauthState.Value {
		c.Redirect(http.StatusTemporaryRedirect, "/")
		helper.GetResponse("invalid oauth google state", 400, true).Send(c)
		return
	}

	// Exchange Auth Code for Tokens
	token, err := config.OAuthConfig.GoogleLoginConfig.Exchange(
		context.Background(), code)

	// ERROR : Auth Code Exchange Failed
	if err != nil {
		helper.GetResponse(err.Error(), 400, true).Send(c)
		return
	}

	// Fetch User Data from google server
	res, err := helper.GetGoogleUser(token)

	// ERROR : Unable to get user data from google
	if err != nil {
		helper.GetResponse(err.Error(), 400, true).Send(c)
		return
	}

	h.loginGoogleUser.Execute(res).Send(c)
}
