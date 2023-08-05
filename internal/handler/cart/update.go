package cart

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *CartHandler) Update(ctx *gin.Context) {
	userId := uint64(1)
	productId, _ := strconv.ParseUint(ctx.Param("productId"), 10, 64)

	h.updateCartUc.Execute(userId, productId).Send(ctx)
}
