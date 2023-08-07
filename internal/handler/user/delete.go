package user

import (
	"github.com/gin-gonic/gin"
)

func (h *userHandler) delete(c *gin.Context) {
	userId := c.MustGet("userId").(uint64)
	h.deleteUser.Execute(userId).Send(c)
}
