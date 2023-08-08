package cart

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/helper"
)

func (ct *CartHandler) post(ctx *gin.Context) {
	userId := ctx.MustGet("userId").(uint64)

	productId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		helper.GetResponse(err.Error(), http.StatusInternalServerError, true)
		return
	}
	ct.insertCartUc.Execute(userId, productId).Send(ctx)
}
