package review

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/helper"
)

func (h *reviewHandler) get(c *gin.Context) {
	productId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		helper.GetResponse(err.Error(), http.StatusInternalServerError, true)
	}
	h.getAllReview.Execute(productId).Send(c)
}
