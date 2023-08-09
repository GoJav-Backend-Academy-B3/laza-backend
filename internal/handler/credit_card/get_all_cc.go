package credit_card

import "github.com/gin-gonic/gin"

// Get All Credit By User Card godoc
// @Summary Get Credit Card by User
// @Description Get All Credit Card
// @Tags creditcard
// @Accept json
// @Produce json
// @Security JWT
// @Success 200 {object} helper.Response{status=string,isError=bool,data=[]response.CreditCardResponse}
// @Failure 500 {object} helper.Response{status=string,description=string,isError=bool}
// @Router /credit-card [GET]
func (h *getCreditCardHandler) GetAll(c *gin.Context) {
	userId := c.MustGet("userId").(uint64)
	h.getAllCcUc.Execute(userId).Send(c)
}
