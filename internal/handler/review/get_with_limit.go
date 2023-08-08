package review

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/helper"
)

func (h *reviewHandler) getWithLimit(c *gin.Context) {
	productId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		helper.GetResponse(err.Error(), http.StatusInternalServerError, true)
	}
	page := c.Query("page")
	perpage := c.Query("perpage")

	pageParse, _ := strconv.ParseUint(page, 10, 64)
	perpageParse, _ := strconv.ParseUint(perpage, 10, 64)

	h.getWithLimitReview.Execute(pageParse, perpageParse, productId).Send(c)
}
