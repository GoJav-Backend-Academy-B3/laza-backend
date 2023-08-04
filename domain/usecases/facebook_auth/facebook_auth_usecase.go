package facebook_auth

import "github.com/gin-gonic/gin"

type FacebookAuthUsecase interface {
	Execute(c *gin.Context) (err error)
}
