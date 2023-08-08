package cart

import "github.com/gin-gonic/gin"

func (h *CartHandler) GetCartOrder(ctx *gin.Context) {
	//userId := ctx.MustGet("userId").(uint64)
	userId := uint64(1)
	h.getCartOrderUc.Execute(userId).Send(ctx)
}
