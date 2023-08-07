package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/helper"
)

func RoleMiddleware(role bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		isAllowed := c.MustGet("role").(bool)

		if isAllowed != role {
			helper.GetResponse("you do not have permission to access", 401, true).Send(c)
			c.Abort()
			return
		}

		c.Set("role", isAllowed)
		c.Next()
	}
}
