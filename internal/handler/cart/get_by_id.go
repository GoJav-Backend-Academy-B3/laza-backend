package cart

import "github.com/gin-gonic/gin"

func (h *CartHandler) GetById(ctx *gin.Context) {
	//userId := ctx.MustGet("userId").(uint64)

	userId := uint64(1)
	h.getCartByIdUc.Execute(userId).Send(ctx)
}
