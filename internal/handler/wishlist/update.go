package wishlist

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/helper"
)

// Update Wishlist godoc
// @Summary Update Wishlist
// @Description Update wish Wishlist
// @Tags wishlist
// @Accept json
// @Produce json
// @Param id path int true "ID of the product"
// @Security JWT
// @Success 200 {object} helper.Response{status=string,isError=bool,data=string}
// @Success 404 {object} helper.Response{status=string,isError=bool,data=string}
// @Failure 500 {object} helper.Response{status=string,description=string,isError=bool}
// @Router /products/{id}/wishlists [PUT]
func (h *getWishlistHandler) Put(ctx *gin.Context) {
	userId := ctx.MustGet("userId").(uint64)

	productid, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		helper.GetResponse(err.Error(), http.StatusInternalServerError, true)
		return
	}

	value, err := h.updateWishlistUsecase.Execute(userId, productid)

	if err != nil {
		helper.GetResponse(err.Error(), 500, true).Send(ctx)
		return
	}

	helper.GetResponse(value, 200, false).Send(ctx)

}
