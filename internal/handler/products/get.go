package products

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/defaults"
	"github.com/phincon-backend/laza/domain/response"
	"github.com/phincon-backend/laza/helper"
)

// GetProducts godoc
// @Summary Gets all products
// @Description Gets all products with the abilities to serch queries
// @Tags product
// @Produce json
// @Param limit query int false "limit for pagination"
// @Param  offset query int false "offset for pagination"
// @Param search query string false "Search keyword for products"
// @Success 200 {object} helper.Response{code=string,isError=bool,status=string,data=[]response.ProductOverview}
// @Failure 500 {object} helper.Response{code=int,description=string,isError=bool}
// @Router /products [get]
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

	if len(search_q) == 0 {
		products, err := h.viewProductUsecase.Execute(offset, limit)
		if err != nil {
			response := helper.GetResponse(err, http.StatusInternalServerError, 1 == 1)
			c.JSON(http.StatusInternalServerError, response)
			return
		}

		productsResponse := make([]response.ProductOverview, 0)
		for _, each := range products {
			var product response.ProductOverview
			product.FillFromEntity(each)
			productsResponse = append(productsResponse, product)
		}
		response := helper.GetResponse(productsResponse, http.StatusOK, false)
		response.Send(c)
	} else {
		products, err := h.searchProductByNameUsecase.Execute(search_q, offset, limit)
		if err != nil {
			response := helper.GetResponse(err, http.StatusInternalServerError, 1 == 1)
			c.JSON(http.StatusInternalServerError, response)
			return
		}

		productsResponse := make([]response.ProductOverview, 0)
		for _, each := range products {
			var product response.ProductOverview
			product.FillFromEntity(each)
			productsResponse = append(productsResponse, product)
		}
		response := helper.GetResponse(productsResponse, 200, false)

		c.JSON(response.Code, response)
		return
	}
}
