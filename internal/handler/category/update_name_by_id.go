package category

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/domain/request"
	"github.com/phincon-backend/laza/helper"
	"gorm.io/gorm"
	"net/http"
)

func (ch *categoryHandler) updateNameById(c *gin.Context) {
	var categoryRequest request.CategoryRequest

	if err := c.ShouldBindJSON(&categoryRequest); err != nil {
		helper.GetResponse(err.Error(), http.StatusBadRequest, true).Send(c)
		return
	}

	result, err := ch.updateCategoryNameByIdUsecase.Execute(categoryRequest)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.GetResponse(err.Error(), http.StatusNotFound, true).Send(c)
			return
		}
		helper.GetResponse(err.Error(), http.StatusBadRequest, true).Send(c)
		return
	}

	helper.GetResponse(result, http.StatusOK, false).Send(c)
	return
}
