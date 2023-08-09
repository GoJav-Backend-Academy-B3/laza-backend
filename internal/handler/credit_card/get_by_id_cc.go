package credit_card

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/helper"
)

func (h *getCreditCardHandler) GetById(c *gin.Context) {
	ccId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		helper.GetResponse(err.Error(), http.StatusInternalServerError, true)
		return
	}
	h.getByIdCcUc.Execute(ccId).Send(c)
}
