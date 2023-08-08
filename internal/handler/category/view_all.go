package category

import (
	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/helper"
	"net/http"
)

func (ch *categoryHandler) getAll(c *gin.Context) {
	results, err := ch.viewCategoryUsecase.Execute()
	if err != nil {
		helper.GetResponse(err.Error(), http.StatusInternalServerError, true).Send(c)
		return
	}
	helper.GetResponse(results, http.StatusOK, false).Send(c)
	return
}
