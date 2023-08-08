package products

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/phincon-backend/laza/domain/request"
	"github.com/phincon-backend/laza/domain/response"
	"github.com/phincon-backend/laza/helper"
)

func (h *productHandler) put(c *gin.Context) {

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

	request := request.ProductRequest{}
	err = c.ShouldBindWith(&request, binding.FormMultipart)
	if err != nil {
		helper.GetResponse(
			gin.H{"reason": "Cannot process data", "err": err.Error()},
			http.StatusUnprocessableEntity,
			true).Send(c)
		return
	}

	model, err := h.updateProductUsecase.Execute(id, request)
	if err != nil {
		helper.GetResponse(err.Error(), http.StatusBadRequest, true).Send(c)
		return
	}
	response := response.Product{}
	response.FillFromEntity(model)
	helper.GetResponse(response, http.StatusOK, false).Send(c)
}
