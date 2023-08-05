package cart

import "github.com/gin-gonic/gin"

func (h *CartHandler) GetById(ctx *gin.Context) {
	userId := uint64(1)
	h.getCartByIdUc.Execute(userId).Send(ctx)
}
