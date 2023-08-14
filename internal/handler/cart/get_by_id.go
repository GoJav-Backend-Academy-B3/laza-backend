package cart

import (
	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/helper"
)

// Get Cart godoc
// @Summary Get Cart
// @Description Get cart
// @Tags cart
// @Accept json
// @Produce json
// @Security JWT
// @Success 200 {object} helper.Response{status=string,isError=bool,data=response.CartInfo}
// @Failure 500 {object} helper.Response{status=string,description=string,isError=bool}
// @Router /carts [GET]
func (h *CartHandler) GetById(ctx *gin.Context) {
	userId := ctx.MustGet("userId").(uint64)

	_result, err := h.getCartByIdUc.Execute(userId)
	if err != nil {
		helper.GetResponse(err.Error(), 500, true).Send(ctx)
		return
	}

	helper.GetResponse(_result, 200, false).Send(ctx)
}
