package size

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/request"
	"github.com/phincon-backend/laza/domain/response"
	"github.com/phincon-backend/laza/helper"
)

func (h *sizeHandler) post(c *gin.Context) {

	request := request.POSTSizeRequest{}
	err := c.ShouldBindJSON(&request)
	if err != nil {
		helper.GetResponse(
			gin.H{"reason": "Cannot process data", "err": err.Error()},
			http.StatusUnprocessableEntity,
			true).Send(c)
		return
	}

	sizeModels := make([]model.Size, 0)
	for _, v := range request.Sizes {
		size, err := h.addSizeUsecase.Execute(v)
		if err != nil {
			helper.GetResponse(
				gin.H{
					"err":    fmt.Sprintf("Tidak dapat memasukkan size %v", v),
					"reason": err.Error(),
				}, http.StatusUnprocessableEntity, true,
			)
		}
		sizeModels = append(sizeModels, size)
	}
	sizesResponse := make([]response.Size, 0)
	for _, v := range sizeModels {
		var size response.Size
		size.FillFromEntity(v)
		sizesResponse = append(sizesResponse, size)
	}
	helper.GetResponse(sizesResponse, http.StatusCreated, false).Send(c)
}
