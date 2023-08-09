package facebook_auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/config"
	"github.com/phincon-backend/laza/helper"
)

// FBLogin godoc
// @Summary Facebook login oauth
// @Description this endpoint is only used to generate redirect URL to facebook login dialog.
// @Tags facebookauth
// @Accept json
// @Produce json
// @Failure 405 {object} helper.Response{code=int,description=string,isError=bool}
// @Success 307 {string} string "Redirecting..."
// @Router /auth/facebook [get]
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
	redirectURL := config.OAuthConfig.FacebookLoginConfig.AuthCodeURL(oauthState)
	c.Redirect(http.StatusTemporaryRedirect, redirectURL)
}
