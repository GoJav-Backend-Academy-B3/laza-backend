package size

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/domain/response"
	"github.com/phincon-backend/laza/helper"
)

func (h *sizeHandler) get(c *gin.Context) {

	// Get query id
	id_q := c.Query(QUERY_ID)

	if len(id_q) > 0 {
		id, err := strconv.ParseUint(id_q, 10, 32)
		if err != nil {
			helper.GetResponse(gin.H{
				"error": "id tidak valid",
			}, http.StatusBadRequest, true).Send(c)
			return
		}
		// Call getByIdH passing gin context with id
		h.getByIdH(c)(id)
		return
	} else {
		h.getAllH(c)()
	}
}

func (h *sizeHandler) getByIdH(c *gin.Context) func(id uint64) {
	return func(id uint64) {
		model, err := h.getSizeByIdUsecase.Execute(id)
		if err != nil {
			response := helper.GetResponse(err.Error(), http.StatusInternalServerError, 1 == 1)
			c.JSON(http.StatusInternalServerError, response)
			return
		}

		var size response.Size
		size.FillFromEntity(model)
		helper.GetResponse(size, http.StatusOK, false).Send(c)
	}
}

func (h *sizeHandler) getAllH(c *gin.Context) func() {
	return func() {
		sizes, err := h.getAllSizeUsecase.Execute()
		if err != nil {
			response := helper.GetResponse(err, http.StatusInternalServerError, 1 == 1)
			c.JSON(http.StatusInternalServerError, response)
			return
		}

		sizesResponse := make([]response.Size, 0)
		for _, each := range sizes {
			var size response.Size
			size.FillFromEntity(each)
			sizesResponse = append(sizesResponse, size)
		}
		helper.GetResponse(sizesResponse, http.StatusOK, false).Send(c)
	}
}
