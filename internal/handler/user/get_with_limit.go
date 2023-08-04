package user

import (
	"github.com/gin-gonic/gin"
)

func (h *userHandler) getWithLimit(c *gin.Context) {
	page := c.Query("page")
	perpage := c.Query("perpage")

	h.getWithLimitUser.Execute(uint64(c.GetInt64(perpage)), uint64(c.GetInt64(page))).Send(c)
}
