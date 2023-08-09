package products

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/phincon-backend/laza/domain/requests"
	"github.com/phincon-backend/laza/domain/response"
	"github.com/phincon-backend/laza/helper"
)

// EditProduct godoc
// @Summary Updates product to database
// @Description Updates product to database with form-data as input
// @Tags product
// @Accept mpfd
// @Produce json
// @Param name formData string true "Product name"
// @Param description formData string true "Product description"
// @Param image formData file true "Product images"
// @Param price formData number true "Product price"
// @Param brand formData string true "Product brand (must exists in database)"
// @Param category formData string true "Product category (must exists in database)"
// @Param id path int true "ID of the product"
// TODO: Swagger not separating form value for []string
// @Param sizes formData []string true "Product available sizes"
// @Security JWT
// @Success 200 {object} helper.Response{isError=bool,status=string,data=response.Product}
// @Failure 422 {object} helper.Response{isError=bool,status=string,description=map[string]string}
// @Failure 400 {object} helper.Response{isError=bool,status=string,description=map[string]string}
// @Router /products/{id} [put]
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

	request := requests.ProductRequest{}
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
