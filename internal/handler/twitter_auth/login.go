package twitterauth

import (
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
	"github.com/phincon-backend/laza/domain/response"
	"github.com/phincon-backend/laza/helper"
)

type ProviderIndex struct {
	Providers    []string
	ProvidersMap map[string]string
}

// TwitterLogin godoc
// @Summary Twitter Login auth
// @Description This endpoint is login endpoint for twitter
// @Tags twitterauth
// @Accept json
// @Produce json
// @Failure 405 {object} helper.Response{code=int,description=string,isError=bool}
// @Success 307 {string} string "Redirecting..."
// @Router /auth [get]
func (h *twitterAuthHandler) loginTwitter(c *gin.Context) {

	gothic.Store = helper.GetStore()

	if gothUser, err := gothic.CompleteUserAuth(c.Writer, c.Request); err == nil {
		rb := response.FillFromTwitter(gothUser.Email, gothUser.Name, gothUser.NickName, gothUser.RawData["profile_image_url_https"].(string))
		h.useCaseTwitter.Execute(rb).Send(c)
	} else {
		gothic.BeginAuthHandler(c.Writer, c.Request)
	}
}
