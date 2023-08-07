package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/helper"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var authHeader string

		if authHeader = c.GetHeader("X-Auth-Token"); authHeader == "" {
			helper.GetResponse("header not provide, please login", 401, true).Send(c)
			c.Abort()
			return
		}

		if !strings.Contains(authHeader, "Bearer") {
			helper.GetResponse("invalid header type", 401, true).Send(c)
			c.Abort()
			return
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", -1)

		validatedToken, err := helper.VerifyToken(tokenString)
		if err != nil {
			helper.GetResponse(err.Error(), 401, true).Send(c)
			c.Abort()
			return
		}

		c.Set("authID", validatedToken.UserId)
		c.Next()
	}
}
