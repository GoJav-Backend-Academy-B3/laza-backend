package user

import "github.com/gin-gonic/gin"

func (h *userHandler) get(c *gin.Context) {
	h.getAllUser.Execute().Send(c)
}