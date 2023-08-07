package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/config"
	"github.com/phincon-backend/laza/helper"
)

func (h *authHandler) loginGoogle(c *gin.Context) {
	// Create oauthState cookie
	oauthState := helper.GenerateStateOauthCookie(c)
	/*
		AuthCodeURL receive state that is a token to protect the user
		from CSRF attacks. You must always provide a non-empty string
		and validate that it matches the the state query parameter
		on your redirect callback.
	*/
	u := config.AppsConfig.GoogleLoginConfig.AuthCodeURL(oauthState)

	c.Redirect(http.StatusTemporaryRedirect, u)
}
