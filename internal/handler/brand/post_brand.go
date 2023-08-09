package brand

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/phincon-backend/laza/domain/requests"
	"github.com/phincon-backend/laza/helper"
)

// CreateBrand godoc
// @Summary Post Details
// @Description Post details of brand
// @Tags brand
// @Accept multipart/form-data
// @Produce json
// @Param brand formData requests.BrandRequest true "create brand"
// @Param logo_url formData file true "brand logo"
// @Security JWT
// @Success 201 {object} helper.Response{code=string,isError=bool,status=string,data=model.Brand}
// @Failure 400 {object} helper.Response{code=int,description=string,isError=bool}
// @Failure 500 {object} helper.Response{code=int,description=string,isError=bool}
// @Router /brand [post]
func (h *brandHandler) PostBrandHandler(ctx *gin.Context) {
	var request requests.BrandRequest

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

	brand, err := h.createBrandUsecase.Execute(request)
	if err != nil {
		helper.GetResponse(err.Error(), http.StatusBadRequest, true).Send(ctx)
		return
	}

	helper.GetResponse(brand, http.StatusCreated, false).Send(ctx)
}
