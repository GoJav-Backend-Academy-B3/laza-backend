package wishlist

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/helper"
)

func (h *getWishlistHandler) Put(ctx *gin.Context) {
	userId := ctx.MustGet("userId").(uint64)

	productid, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		helper.GetResponse(err.Error(), http.StatusInternalServerError, true)
		return
	}

	h.updateWishlistUsecase.Execute(userId, productid).Send(ctx)

}
