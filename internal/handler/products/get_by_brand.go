package products

import (
	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/defaults"
	"github.com/phincon-backend/laza/domain/response"
	"github.com/phincon-backend/laza/helper"
	"github.com/phincon-backend/laza/mapper"
	"net/http"
	"strconv"
)

// FindProductByBrand godoc
// @Summary Find product by brand
// @Description It returns a list of products of a brand
// @Tags product
// @Accept json
// @Produce json
// @Security JWT
// @Param name query string true "Brand name"
// @Param limit query int false "limit for pagination"
// @Param  offset query int false "offset for pagination"
// @Success 200 {object} helper.Response{code=string,isError=bool,status=string,data=[]response.Product}
// @Failure 500 {object} helper.Response{code=int,description=string,isError=bool}
// @Router /products/brand [get]
func (pb *viewProductByBrandHandler) get(c *gin.Context) {

	// Get limit and offset query string
	limit_q := c.DefaultQuery(QUERY_LIMIT, defaults.Query.LimitString())
	offset_q := c.DefaultQuery(QUERY_OFFSET, defaults.Query.OffsetString())
	brand := c.Query(QUERY_NAME)

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
		helper.GetResponse(err, http.StatusInternalServerError, true).Send(c)
		return
	}

	productsResponse := make([]response.Product, 0)
	for _, item := range products {
		productsResponse = append(productsResponse, mapper.ProductModelToProductResponse(item))
	}
	helper.GetResponse(productsResponse, 200, false).Send(c)
	return

}
