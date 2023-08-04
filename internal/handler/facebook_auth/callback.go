package facebook_auth

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/config"
	"github.com/phincon-backend/laza/domain/response"
	"io"
	"net/http"
	"strings"
)

func (fb *facebookAuthHandler) FbCallback(c *gin.Context) {
	// check is method is correct
	if c.Request.Method != http.MethodGet {
		response.GetResponse("method not allowed", http.StatusMethodNotAllowed, true).Send(c)
		return
	}

	// get oauth state from cookie for this user
	oauthState, _ := c.Cookie("oauthstate")
	state := c.Request.FormValue("oauthstate")
	code := c.Request.FormValue("code")

	// ERROR : Invalid OAuth State
	if !strings.EqualFold(state, oauthState) {
		c.Redirect(http.StatusTemporaryRedirect, "/")
		fmt.Fprintf(c.Writer, "invalid oauth google state")
		return
	}

	// Exchange Auth Code for Tokens
	token, err := config.FBConfig.LoginConfig.Exchange(
		context.Background(), code)

	// ERROR : Auth Code Exchange Failed
	if err != nil {
		response.GetResponse(fmt.Sprintf("failed code exchange %s", err.Error()), http.StatusUnauthorized, true).Send(c)
		return
	}

	// Fetch User Data from facebook server
	res, err := http.Get(config.OauthFacebookUrlAPI + token.AccessToken)

	// ERROR : Unable to get user data from google
	if err != nil {
		response.GetResponse(fmt.Sprintf("failed code exchange %s", err.Error()), http.StatusUnauthorized, true).Send(c)
		return
	}

	// Parse user data JSON Object
	defer res.Body.Close()
	contents, err := io.ReadAll(res.Body)
	if err != nil {
		response.GetResponse(fmt.Sprintf("failed code exchange %s", err.Error()), http.StatusUnauthorized, true).Send(c)
		return
	}

	// send back response to browser
	response.GetResponse(string(contents), http.StatusOK, false).Send(c)
}
