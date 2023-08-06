package cart

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func (ct *CartHandler) post(ctx *gin.Context) {
	// userId, bl := ctx.Get("authID")
	userId := uint64(1)
	productId, _ := strconv.ParseUint(ctx.Param("productId"), 10, 64)

	ct.insertCartUc.Execute(userId, productId).Send(ctx)
}
