package products

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/phincon-backend/laza/domain/request"
	"github.com/phincon-backend/laza/domain/response"
	"github.com/phincon-backend/laza/helper"
)

func (h *productHandler) post(c *gin.Context) {
	request := request.ProductRequest{}
	err := c.ShouldBindWith(&request, binding.FormMultipart)
	if err != nil {
		helper.GetResponse(
			gin.H{"reason": "Cannot process data", "err": err.Error()},
			http.StatusUnprocessableEntity,
			true).Send(c)
		return
	}
	log.Printf("%v: size: %v\n", request.Image.Filename, request.Image.Size)

	model, err := h.createProductUsecase.Execute(request)
	if err != nil {
		helper.GetResponse(err.Error(), http.StatusBadRequest, true).Send(c)
		return
	}
	response := response.Product{}
	response.FillFromEntity(model)
	helper.GetResponse(response, http.StatusCreated, false).Send(c)
}
