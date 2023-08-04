package wishlist

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/helper"
)

func (h *getWishlistHandler) Put(ctx *gin.Context) {
	ws := model.Wishlist{}

	userid := uint64(1)

	productid, err := strconv.ParseUint(ctx.Param("productId"), 10, 32)
	if err != nil {
		response := helper.GetResponse(err, http.StatusBadRequest, true)
		response.Send(ctx)
	}

	ws.UserId, ws.ProductId = userid, productid
	rs, err := h.updateWishlistUsecase.Execute(ws)
	if err != nil {
		response := helper.GetResponse(err, http.StatusInternalServerError, true)
		response.Send(ctx)
	}

	response := helper.GetResponse(rs, http.StatusOK, false)
	response.Send(ctx)
}