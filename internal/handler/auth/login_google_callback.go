package auth

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/config"
	"github.com/phincon-backend/laza/helper"
)

// GoogleCallback godoc
// @Summary Google callback endpoint when login success.
// @Description This endpoint is a callback endpoint for google login or sign up.
// @Tags googleauth
// @Accept json
// @Produce json
// @Success 200 {object} helper.Response{code=string,isError=bool,status=string,data=map[string]string}
// @Failure 405 {object} helper.Response{description=string,code=int,isError=bool}
// @Failure 403 {object} helper.Response{description=string,code=int,isError=bool}
// @Router /login-google/callback [get]
func (h *authHandler) loginGoogleCallback(c *gin.Context) {
	// get oauth state from cookie for this user
	oauthState, _ := c.Cookie("oauthstate")
	state := c.Query("state")
	code := c.Query("code")

	// ERROR : Invalid OAuth State
	if !strings.EqualFold(state, oauthState) {
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
