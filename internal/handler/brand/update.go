package brand

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/phincon-backend/laza/domain/requests"
	"github.com/phincon-backend/laza/helper"
)

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
