package bank

import "github.com/gin-gonic/gin"

// GetAllBank godoc
// @Summary Get All Bank
// @Description Get all Bank
// @Tags bank
// @Accept json
// @Produce json
// @Security JWT
// @Success 200 {object} helper.Response{code=string,isError=bool,status=string,data=model.Bank}
// @Failure 400 {object} helper.Response{code=int,description=string,isError=bool}
// @Failure 500 {object} helper.Response{code=int,description=string,isError=bool}
// @Router /bank [get]
func (h *bankHandler) get(c *gin.Context) {
	h.getAllBank.Execute().Send(c)
}
