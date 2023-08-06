package cart

import "github.com/gin-gonic/gin"

func (h *CartHandler) GetCartOrder(ctx *gin.Context) {
	// userId, bl := ctx.Get("authID")
	userId := uint64(1)
	h.getCartOrderUc.Execute(userId).Send(ctx)
}
