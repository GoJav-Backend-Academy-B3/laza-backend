package products

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/helper"
)

// GetProductById godoc
// @Summary Get product identified by the given id
// @Description get the product corresponding to the input Id
// @Tags Product
// @Accept json
// @Produce json
// @Param id path int true "ID of the Product"
// @Security JWT
// @Success 200 {object} helper.Response{isError=bool,status=string,data=response.ProductDetail}
// @Router /products/{id} [get]
func (h *productHandler) getById(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		helper.GetResponse("missing 'id' path params", 400, true).Send(c)
		return
	}

	idParse, _ := strconv.ParseUint(id, 10, 64)
	h.getByIdProduct.Execute(idParse).Send(c)
}
