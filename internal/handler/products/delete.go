package products

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/helper"
)

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
