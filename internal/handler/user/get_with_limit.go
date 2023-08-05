package user

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *userHandler) getWithLimit(c *gin.Context) {
	page := c.Query("page")
	perpage := c.Query("perpage")

	pageParse, _ := strconv.ParseUint(page, 10, 64)
	perpageParse, _ := strconv.ParseUint(perpage, 10, 64)

	h.getWithLimitUser.Execute(pageParse, perpageParse).Send(c)
}
