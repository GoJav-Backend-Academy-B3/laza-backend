package brand

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/phincon-backend/laza/domain/requests"
	"github.com/phincon-backend/laza/helper"
)

// UpdateBrand godoc
// @Summary put Details
// @Description put details of brand
// @Tags brand
// @Accept multiprat/form-data
// @Produce json
// @Param id path int true "ID of the brand"
// @Param brand formData requests.BrandRequest true "create brand"
// @Param logo_url formData file true "brand logo"
// @Security JWT
// @Success 200 {object} helper.Response{code=string,isError=bool,status=string,data=model.Brand}
// @Failure 400 {object} helper.Response{code=int,description=string,isError=bool}
// @Failure 500 {object} helper.Response{code=int,description=string,isError=bool}
// @Router /brand/{id} [put]
func (h *brandHandler) UpdateBrand(ctx *gin.Context) {
	var request requests.BrandRequest

	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)

	err := ctx.ShouldBindWith(&request, binding.FormMultipart)
	if err != nil {
		helper.GetResponse(err.Error(), http.StatusBadRequest, true).Send(ctx)
		return
	}

	err = h.validate.Struct(request)
	if err != nil {
		helper.GetResponse(err.Error(), http.StatusBadRequest, true).Send(ctx)
		return
	}

	brand, err := h.updateBrandNameByIdUsecase.Execute(id, request)
	if err != nil {
		helper.GetResponse(err.Error(), http.StatusBadRequest, true).Send(ctx)
		return
	}

	helper.GetResponse(brand, http.StatusOK, false).Send(ctx)
}
