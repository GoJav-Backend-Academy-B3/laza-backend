package wishlist

import "github.com/gin-gonic/gin"

func (h *getWishlistHandler) get(ctx *gin.Context) {
	userId := ctx.MustGet("userId").(uint64)

	h.getWishlistUsecase.Execute(userId).Send(ctx)
}
