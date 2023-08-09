package brand

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/helper"
	"gorm.io/gorm"
)

// GetByBrandId godoc
// @Summary Get Details for a given id
// @Description Get details of brand corresponding is the input Id
// @Tags brand
// @Accept json
// @Produce json
// @Param id path int true "ID of the brand"
// @Success 200 {object} helper.Response{code=string,isError=bool,status=string,data=model.Brand}
// @Failure 400 {object} helper.Response{code=int,description=string,isError=bool}
// @Failure 500 {object} helper.Response{code=int,description=string,isError=bool}
// @Router /brand/{id} [get]
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
