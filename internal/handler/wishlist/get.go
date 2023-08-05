package wishlist

import "github.com/gin-gonic/gin"

func (h *getWishlistHandler) get(ctx *gin.Context) {
	userId := uint64(1)
	h.getWishlistUsecase.Execute(userId).Send(ctx)
}
