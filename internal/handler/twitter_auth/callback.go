package twitterauth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
	"github.com/phincon-backend/laza/helper"
)

func (h *twitterAuthHandler) TwitterCallBack(c *gin.Context) {
	gothic.Store = helper.GetStore()
	user, err := gothic.CompleteUserAuth(c.Writer, c.Request)

	if err != nil {
		helper.GetResponse(err.Error(), http.StatusUnauthorized, true).Send(c)
		return
	}

	// send back response to browse
	helper.GetResponse(user, http.StatusOK, false).Send(c)
}
