package products

import (
	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/defaults"
	"github.com/phincon-backend/laza/domain/response"
	"github.com/phincon-backend/laza/helper"
	"net/http"
	"strconv"
)

func (pb *viewProductByBrandHandler) get(c *gin.Context) {

	// Get limit and offset query string
	limit_q := c.DefaultQuery(QUERY_LIMIT, defaults.Query.LimitString())
	offset_q := c.DefaultQuery(QUERY_OFFSET, defaults.Query.OffsetString())
	brand := c.Query(QUERY_BRAND)

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

	products, err := pb.viewProductByBrandUsecase.Execute(brand, offset, limit)
	if err != nil {
		response := helper.GetResponse(err, http.StatusInternalServerError, 1 == 1)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	productsResponse := make([]response.Product, 0)
	for _, each := range products {
		var product response.Product
		product.FillFromEntity(each)
		productsResponse = append(productsResponse, product)
	}
	results := helper.GetResponse(productsResponse, 200, false)

	c.JSON(results.Code, results)
	return

}
