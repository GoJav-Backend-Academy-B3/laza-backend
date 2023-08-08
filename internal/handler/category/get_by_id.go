package category

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/helper"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func (ch *categoryHandler) getById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helper.GetResponse(err.Error(), http.StatusBadRequest, true).Send(c)
		return
	}

	result, err := ch.getCategoryByIdUsecase.Execute(uint64(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.GetResponse(err.Error(), http.StatusNotFound, true).Send(c)
			return
		}
		helper.GetResponse(err.Error(), http.StatusInternalServerError, true).Send(c)
		return
	}

	helper.GetResponse(result, http.StatusOK, false).Send(c)
	return
}
