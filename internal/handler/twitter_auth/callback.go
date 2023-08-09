package twitterauth

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
	"github.com/phincon-backend/laza/domain/response"
	"github.com/phincon-backend/laza/helper"
)

func (h *twitterAuthHandler) twitterCallBack(c *gin.Context) {
	// gothic.Store = helper.GetStore()
	gothUser, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	if err != nil {
		fmt.Println(gothUser)
		helper.GetResponse(err.Error(), http.StatusUnauthorized, true).Send(c)
		return
	}

	rb := response.FillFromTwitter(gothUser.Email, gothUser.Name, gothUser.NickName, gothUser.RawData["profile_image_url_https"].(string))

	h.useCaseTwitter.Execute(rb).Send(c)
}
