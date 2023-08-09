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

	user, err := gothic.CompleteUserAuth(c.Writer, c.Request)

	if err != nil {
		fmt.Println(user)
		helper.GetResponse(err.Error(), http.StatusUnauthorized, true).Send(c)
		return
	}

	rb := response.FillFromTwitter(user.Email, user.Name, user.NickName, user.RawData["profile_image_url_https"].(string))

	h.useCaseTwitter.Execute(rb).Send(c)
}
