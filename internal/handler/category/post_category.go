package category

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/domain/requests"
	"github.com/phincon-backend/laza/helper"
)

func (ch *categoryHandler) postCategory(c *gin.Context) {
	var categoryRequest requests.CategoryRequest
	if err := c.ShouldBindJSON(&categoryRequest); err != nil {
		helper.GetResponse(err.Error(), http.StatusBadRequest, true).Send(c)
		return
	}
	log.Println(categoryRequest)
	result, err := ch.createCategoryUsecase.Execute(categoryRequest)
	if err != nil {
		helper.GetResponse(err.Error(), http.StatusInternalServerError, true).Send(c)
		return
	}

	helper.GetResponse(result, http.StatusCreated, false).Send(c)
	return
}
