package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/helper"
)

func AdminRoleMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		isAllowed := c.MustGet("userRole").(bool)

		if !isAllowed {
			helper.GetResponse("you do not have permission to access", 401, true).Send(c)
			c.Abort()
			return
		}

		c.Next()
	}
}
