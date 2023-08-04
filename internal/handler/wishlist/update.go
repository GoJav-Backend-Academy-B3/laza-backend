package wishlist

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/response"
)

func (h *getWishlistHandler) Put(ctx *gin.Context) {
	ws := model.Wishlist{}
	userid, err := strconv.ParseUint(ctx.Param("userId"), 10, 32)
	if err != nil {
		response := response.GetResponse(err, http.StatusBadRequest, true)
		response.Send(ctx)
	}
	productid, err := strconv.ParseUint(ctx.Param("productId"), 10, 32)
	if err != nil {
		response := response.GetResponse(err, http.StatusBadRequest, true)
		response.Send(ctx)
	}

	ws.UserId, ws.ProductId = userid, productid
	rs, err := h.updateWishlistUsecase.Execute(ws)
	if err != nil {
		response := response.GetResponse(err, http.StatusInternalServerError, true)
		response.Send(ctx)
	}

	fmt.Println(rs)
	response := response.GetResponse(rs, http.StatusOK, false)
	response.Send(ctx)
}
