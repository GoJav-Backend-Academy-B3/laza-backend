package category

import (
	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/domain/request"
	"github.com/phincon-backend/laza/helper"
	"log"
	"net/http"
)

func (ch *categoryHandler) postCategory(c *gin.Context) {
	var categoryRequest request.CategoryRequest
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
