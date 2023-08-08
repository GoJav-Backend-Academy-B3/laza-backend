package bank

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/phincon-backend/laza/domain/requests"
	"github.com/phincon-backend/laza/helper"
)

// CreateBank godoc
// @Summary Post Details
// @Description Post details of bank
// @Tags bank
// @Accept multipart/form-data
// @Produce json
// @Param bank formData requests.BankRequest true "create bank"
// @Param image formData file true "bank"
// @Security JWT
// @Success 201 {object} helper.Response{code=string,isError=bool,status=string,data=model.Bank}
// @Failure 400 {object} helper.Response{code=int,description=string,isError=bool}
// @Failure 500 {object} helper.Response{code=int,description=string,isError=bool}
// @Router /bank [post]
func (h *bankHandler) insert(c *gin.Context) {
	var request requests.BankRequest
	if err := c.ShouldBindWith(&request, binding.FormMultipart); err != nil {
		helper.GetResponse(err.Error(), 400, true).Send(c)
		return
	}

	err := h.validate.Struct(request)
	if err != nil {
		helper.GetResponse(err.Error(), http.StatusBadRequest, true).Send(c)
		return
	}

	h.insertBank.Execute(request).Send(c)
}
