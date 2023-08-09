package brand

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/helper"
)

func (h *brandHandler) ViewAllBrand(ctx *gin.Context) {
	result, err := h.viewBrandUsecase.Execute()
	if err != nil {
		helper.GetResponse(err.Error(), http.StatusInternalServerError, true).Send(ctx)
		return
	}

	helper.GetResponse(result, http.StatusOK, true).Send(ctx)

}
