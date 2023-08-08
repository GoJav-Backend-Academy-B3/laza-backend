package category

import (
	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/helper"
	"net/http"
)

func (ch *categoryHandler) searchByName(c *gin.Context) {
	name := c.Query("name")
	results, err := ch.searchCategoryByNameUsecase.Execute(name)
	if err != nil {
		helper.GetResponse(err.Error(), http.StatusInternalServerError, true).Send(c)
		return
	}
	helper.GetResponse(results, http.StatusOK, false).Send(c)
	return
}
