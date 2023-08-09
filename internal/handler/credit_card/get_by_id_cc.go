package credit_card

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/helper"
)

// Get Credit Card by Id godoc
// @Summary Get Credit Card By Id
// @Description Get Credit Card
// @Tags creditcard
// @Accept json
// @Produce json
// @Param id path int true "ID of the credit card"
// @Security JWT
// @Success 200 {object} helper.Response{status=string,isError=bool,data=response.CreditCardResponse}
// @Failure 400 {object} helper.Response{status=string,description=string,isError=bool}
// @Failure 500 {object} helper.Response{status=string,description=string,isError=bool}
// @Router /credit-card/{id} [GET]
func (h *getCreditCardHandler) GetById(c *gin.Context) {
	ccId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		helper.GetResponse(err.Error(), http.StatusInternalServerError, true)
		return
	}
	h.getByIdCcUc.Execute(ccId).Send(c)
}
