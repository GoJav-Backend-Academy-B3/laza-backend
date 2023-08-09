package cart

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/helper"
)

// Delete Cart
// @Summary Delete Cart
// @Description Delete cart with product id
// @Tags cart
// @Accept json
// @Produce json
// @Param id path int true "ID of the product"
// @Security JWT
// @Success 200 {object} helper.Response{status=string,isError=bool,data=string}
// @Failure 401 {object} helper.Response{status=string,description=string,isError=bool}
// @Failure 500 {object} helper.Response{status=string,description=string,isError=bool}
// @Router /products/{id}/carts [delete]
func (ct *CartHandler) Delete(ctx *gin.Context) {
	userId := ctx.MustGet("userId").(uint64)
	productId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		helper.GetResponse(err.Error(), http.StatusInternalServerError, true)
		return
	}

	ct.deleteCartUc.Execute(userId, productId).Send(ctx)
}
