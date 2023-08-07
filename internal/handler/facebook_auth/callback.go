package facebook_auth

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/config"
	"github.com/phincon-backend/laza/domain/response"
	"github.com/phincon-backend/laza/helper"
	"io"
	"net/http"
	"strings"
)

func (fb *facebookAuthHandler) FbCallback(c *gin.Context) {
	// check is method is correct
	if c.Request.Method != http.MethodGet {
		helper.GetResponse("method not allowed", http.StatusMethodNotAllowed, true).Send(c)
		return
	}

	// get oauth state from cookie for this user
	oauthState, _ := c.Cookie("oauthstate")
	state := c.Query("state")
	code := c.Query("code")

	// ERROR : Invalid OAuth State
	if !strings.EqualFold(state, oauthState) {
		fmt.Fprintf(c.Writer, "invalid oauth google state")
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}

	// Exchange Auth Code for Tokens
	token, err := config.OAuthConfig.FacebookLoginConfig.Exchange(
		context.Background(), code)

	// ERROR : Auth Code Exchange Failed
	if err != nil {
		helper.GetResponse(fmt.Sprintf("failed code exchange %s", err.Error()), http.StatusUnauthorized, true).Send(c)
		return
	}

	// Fetch User Data from facebook server
	res, err := http.Get(config.OauthFacebookUrlAPI + token.AccessToken)

	// ERROR : Unable to get user data from google
	if err != nil {
		helper.GetResponse(fmt.Sprintf("failed code exchange %s", err.Error()), http.StatusUnauthorized, true).Send(c)
		return
	}

	// Parse user data JSON Object
	defer res.Body.Close()
	contents, err := io.ReadAll(res.Body)
	if err != nil {
		helper.GetResponse(fmt.Sprintf("failed code exchange %s", err.Error()), http.StatusUnauthorized, true).Send(c)
		return
	}

	var fbResponse response.FBAuthResponse
	if err = json.Unmarshal(contents, &fbResponse); err != nil {
		helper.GetResponse(fmt.Sprintf("failed to get user data %s", err.Error()), http.StatusUnauthorized, true).Send(c)
		return
	}

	accessToken, err := fb.facebookAuthUsecase.Execute(fbResponse)
	if err != nil {
		helper.GetResponse(fmt.Sprintf("failed to create access token %s", err.Error()), http.StatusUnauthorized, true).Send(c)
		return
	}
	responseMap := map[string]string{
		"access_token": accessToken,
	}
	cookie := &http.Cookie{
		Name:     "access_token",
		Value:    accessToken,
		Path:     "/",
		HttpOnly: true,
		MaxAge:   int(config.LoadJWTConfig().GetTokenExpiry()),
	}
	http.SetCookie(c.Writer, cookie)

	// send back response to browse
	helper.GetResponse(responseMap, http.StatusOK, false).Send(c)
}
