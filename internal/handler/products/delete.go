package products

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/helper"
)

// DeleteProduct godoc
// @Summary Deletes product identified by the given id
// @Description Deletes product corresponding to the input Id
// @Tags product
// @Produce json
// @Param id path int true "ID of the product"
// @Security JWT
// @Success 200 {object} helper.Response{code=string,isError=bool,status=string,data=map[string]string}
// @Failure 400 {object} helper.Response{code=int,description=string,isError=bool,description=map[string]string}
// @Failure 500 {object} helper.Response{code=int,description=string,isError=bool,description=map[string]string}
// @Router /products/{id} [delete]
func (h *productHandler) delete(c *gin.Context) {

	id_p := c.Param("id")
	id, err := strconv.ParseUint(id_p, 10, 64)
	if err != nil {
		helper.GetResponse(
			gin.H{"reason": "id tidak valid"},
			http.StatusBadRequest,
			true,
		).Send(c)
		return
	}

	err = h.deleteProductUsecase.Execute(id)
	if err != nil {
		helper.GetResponse(err.Error(), http.StatusBadRequest, false).Send(c)
		return
	}
	helper.GetResponse(gin.H{
		"status": "deleted",
	}, http.StatusOK, false).Send(c)
}
