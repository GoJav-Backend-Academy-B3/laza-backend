package cart

import "github.com/gin-gonic/gin"

// Get Cart godoc
// @Summary Get Cart
// @Description Get cart
// @Tags cart
// @Accept json
// @Produce json
// @Param id path int true "ID of the product"
// @Security JWT
// @Success 200 {object} helper.Response{status=string,isError=bool,data=response.CartInfo}
// @Failure 500 {object} helper.Response{status=string,description=string,isError=bool}
// @Router /products/{id}/carts [GET]
func (h *CartHandler) GetById(ctx *gin.Context) {
	userId := ctx.MustGet("userId").(uint64)
	h.getCartByIdUc.Execute(userId).Send(ctx)
}
