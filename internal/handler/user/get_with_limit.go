package user

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetAllWithLimitUser godoc
// @Summary Get All User With Limit
// @Description Get user all with limit
// @Tags user
// @Accept json
// @Produce json
// @Security JWT
// @Param page query string true "Query Page"
// @Param perpage query string true "Query Perpage "
// @Success 200 {object} helper.Response{code=string,isError=bool,status=string,data=response.User}
// @Failure 400 {object} helper.Response{code=int,description=string,isError=bool}
// @Failure 500 {object} helper.Response{code=int,description=string,isError=bool}
// @Router /user/ [get]
func (h *userHandler) getWithLimit(c *gin.Context) {
	page := c.Query("page")
	perpage := c.Query("perpage")

	pageParse, _ := strconv.ParseUint(page, 10, 64)
	perpageParse, _ := strconv.ParseUint(perpage, 10, 64)

	h.getWithLimitUser.Execute(pageParse, perpageParse).Send(c)
}
