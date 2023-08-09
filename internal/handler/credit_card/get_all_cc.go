package credit_card

import "github.com/gin-gonic/gin"

func (h *getCreditCardHandler) GetAll(c *gin.Context) {
	userId := c.MustGet("userId").(uint64)
	h.getAllCcUc.Execute(userId).Send(c)
}
