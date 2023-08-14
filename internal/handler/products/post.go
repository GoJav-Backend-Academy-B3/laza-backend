package products

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/phincon-backend/laza/domain/requests"
	"github.com/phincon-backend/laza/domain/response"
	"github.com/phincon-backend/laza/helper"
)

// AddProduct godoc
// @Summary Adds product to database
// @Description Adds product to database with form-data as input
// @Tags product
// @Accept mpfd
// @Produce json
// @Param name formData string true "Product name"
// @Param description formData string true "Product description"
// @Param image formData file true "Product images. The file should not exceed 2MiB or approximately equivalent to 2.0971MB"
// @Param price formData number true "Product price"
// @Param brand formData string true "Product brand (must exists in database)"
// @Param category formData string true "Product category (must exists in database)"
// TODO: Swagger not separating form value for []string
// @Param sizes formData []string true "Product available sizes"
// @Security JWT
// @Success 201 {object} helper.Response{isError=bool,status=string,data=response.Product}
// @Failure 422 {object} helper.Response{isError=bool,status=string,description=map[string]string}
// @Failure 400 {object} helper.Response{isError=bool,status=string,description=map[string]string}
// @Router /products [post]
func (h *productHandler) post(c *gin.Context) {
	request := requests.ProductRequest{}
	err := c.ShouldBindWith(&request, binding.FormMultipart)
	if err != nil {
		helper.GetResponse(
			gin.H{"reason": "Cannot process data", "err": err.Error()},
			http.StatusUnprocessableEntity,
			true).Send(c)
		return
	}

	// Check if a file exceeds 2MiB
	if request.Image.Size > (1 << 21) {
		helper.GetResponse(
			"file too big",
			http.StatusRequestEntityTooLarge,
			true).Send(c)
		return
	}

	model, err := h.createProductUsecase.Execute(request)
	if err != nil {
		helper.GetResponse(err.Error(), http.StatusBadRequest, true).Send(c)
		return
	}
	response := response.Product{}
	response.FillFromEntity(model)
	helper.GetResponse(response, http.StatusCreated, false).Send(c)
}
