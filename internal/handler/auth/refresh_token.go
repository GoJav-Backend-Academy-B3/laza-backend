package auth

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/helper"
)

// RefreshToken godoc
// @Summary Refresh Token
// @Description Refresh Token for user
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 {object} helper.Response{code=string,isError=bool,status=string,data=map[string]string}
// @Failure 401 {object} helper.Response{code=int,description=string,isError=bool}
// @Failure 500 {object} helper.Response{code=int,description=string,isError=bool}
// @Router /auth/refresh [get]
func (h *authHandler) refreshToken(c *gin.Context) {
	var authHeader string

	if authHeader = c.GetHeader("X-Auth-Refresh"); authHeader == "" {
		helper.GetResponse("header not provide, please login", 401, true).Send(c)
		return
	}

	if !strings.Contains(authHeader, "Bearer") {
		helper.GetResponse("invalid header type", 401, true).Send(c)
		return
	}

	tokenString := strings.Replace(authHeader, "Bearer ", "", -1)
	validatedToken, err := helper.VerifyRefresh(tokenString)
	if err != nil {
		helper.GetResponse(err.Error(), 401, true).Send(c)
		return
	}

	h.refreshTokenUser.Execute(uint64(validatedToken.UserId)).Send(c)
}
