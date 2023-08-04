package products

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/defaults"
	"github.com/phincon-backend/laza/domain/response"
)

func (h *productHandler) get(c *gin.Context) {

	// Get limit and offset query string
	limit_q := c.DefaultQuery(QUERY_LIMIT, defaults.Query.LimitString())
	offset_q := c.DefaultQuery(QUERY_OFFSET, defaults.Query.OffsetString())
	search_q := c.Query(QUERY_SEARCH)

	// convert limit to unsigned integer
	limit, err := strconv.ParseUint(limit_q, 10, 32)
	// if this fails, set to default value
	if err != nil {
		limit = defaults.Query.Limit()
	}

	// convert offset to unsigned integer
	offset, err := strconv.ParseUint(offset_q, 10, 32)
	// if this fails, set to default value
	if err != nil {
		offset = defaults.Query.Offset()
	}

	products, err := h.viewProductUsecase.Execute(offset, limit)
	productsResponse := make([]response.Product, 0)
	for _, each := range products {
		var product response.Product
		product.FillFromEntity(each)
		productsResponse = append(productsResponse, product)
	}
	response := response.GetResponse(productsResponse, 200, false)

	c.JSON(response.Code, response)
}
