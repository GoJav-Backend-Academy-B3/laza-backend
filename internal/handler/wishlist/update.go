package wishlist

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/phincon-backend/laza/domain/requests"
	"github.com/phincon-backend/laza/helper"
)

// Update Wishlist godoc
// @Summary Update Wishlist
// @Description Update wish Wishlist
// @Tags wishlist
// @Accept json
// @Produce json
// @Param wishlist body requests.WishlistRequest true "add product to wishlist"
// @Security JWT
// @Success 200 {object} helper.Response{status=string,isError=bool,data=string}
// @Failure 500 {object} helper.Response{status=string,description=string,isError=bool}
// @Router /wishlists [PUT]
func (h *getWishlistHandler) Put(ctx *gin.Context) {
	userId := ctx.MustGet("userId").(uint64)

	var requestBody requests.WishlistRequest
	err := ctx.Bind(&requestBody)
	if err != nil {
		helper.GetResponse(err.Error(), 400, true).Send(ctx)
		return
	}

	value, err := h.updateWishlistUsecase.Execute(userId, requestBody)
	if err, ok := err.(validator.ValidationErrors); ok {
		helper.GetResponse(err.Error(), 400, true).Send(ctx)
		return
	}

	if err != nil {
		helper.GetResponse(err.Error(), 500, true).Send(ctx)
		return
	}

	helper.GetResponse(value, 200, false).Send(ctx)

}
