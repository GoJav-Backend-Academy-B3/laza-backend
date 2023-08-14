package twitterauth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
	"github.com/phincon-backend/laza/domain/response"
	"github.com/phincon-backend/laza/helper"
)

// TwitterCallback godoc
// @Summary Twitter callback endpoint when login success
// @Description This endpoint is a callback endpoint when login success
// @Tags twitterauth
// @Accept json
// @Produce json
// @Success 200 {object} helper.Response{code=string,isError=bool,data=map[string]string}
// @Failure 405 {object} helper.Response{code=int,description=string,isError=bool}
// @Failure 403 {object} helper.Response{code=int,description=string,isError=bool}
// @Router /auth/twitter/callback [get]
func (h *twitterAuthHandler) twitterCallBack(c *gin.Context) {

	user, err := gothic.CompleteUserAuth(c.Writer, c.Request)

	if err != nil {
		helper.GetResponse(err.Error(), http.StatusUnauthorized, true).Send(c)
		return
	}

	var responseBody response.TwitterUser
	responseBody.FillEntity(user)

	_result, err := h.useCaseTwitter.Execute(responseBody)
	if err != nil {
		helper.GetResponse(err.Error(), 500, true).Send(c)
		return
	}

	helper.GetResponse(_result, 200, false).Send(c)
}
