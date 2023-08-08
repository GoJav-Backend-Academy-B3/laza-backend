package category

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/helper"
	"net/http"
	"strconv"
)

func (ch *categoryHandler) deleteById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helper.GetResponse(err.Error(), http.StatusBadRequest, true).Send(c)
		return
	}
	rowAffected, err := ch.deleteCategoryByIdUsecase.Execute(uint64(id))
	if err != nil {
		if rowAffected == 0 {
			helper.GetResponse("no record found", http.StatusNotFound, true).Send(c)
			return
		}
		helper.GetResponse(err.Error(), http.StatusInternalServerError, true).Send(c)
		return
	}

	helper.GetResponse(fmt.Sprintf("rows affected %d", rowAffected), http.StatusOK, false).Send(c)
	return

}
