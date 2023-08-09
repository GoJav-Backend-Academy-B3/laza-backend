package brand

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/helper"
)

func (h *brandHandler) SearchByBrandName(ctx *gin.Context) {
	name := ctx.Query("name")

	results, err := h.searchBrandByNameUsecase.Execute(name)
	if err != nil {
		helper.GetResponse(err.Error(), http.StatusInternalServerError, true).Send(ctx)
		return
	}

	helper.GetResponse(results, http.StatusOK, false).Send(ctx)

}
