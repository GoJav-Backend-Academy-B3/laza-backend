package cart

import "github.com/gin-gonic/gin"

func (h *CartHandler) GetById(ctx *gin.Context) {
	// userId, bl := ctx.Get("authID")

	userId := uint64(1)
	h.getCartByIdUc.Execute(userId).Send(ctx)
}
