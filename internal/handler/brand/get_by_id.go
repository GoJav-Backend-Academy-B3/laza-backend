package brand

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/helper"
	"gorm.io/gorm"
)

func (h *brandHandler) GetById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		helper.GetResponse(err.Error(), http.StatusBadRequest, true).Send(ctx)
		return
	}

	result, err := h.getBrandByIdUsecase.Execute(uint64(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.GetResponse(err.Error(), http.StatusNotFound, true).Send(ctx)
			return
		}
		helper.GetResponse(err.Error(), http.StatusInternalServerError, true).Send(ctx)
		return
	}

	helper.GetResponse(result, http.StatusOK, false).Send(ctx)

}
