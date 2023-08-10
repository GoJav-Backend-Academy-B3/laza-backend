package twitterauth

import (
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
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

	gothic.BeginAuthHandler(c.Writer, c.Request)

}
