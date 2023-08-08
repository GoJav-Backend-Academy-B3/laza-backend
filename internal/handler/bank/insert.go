package bank

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/helper"
)

// CreateBank godoc
// @Summary Post Details
// @Description Post details of bank
// @Tags bank
// @Accept multipart/form-data
// @Produce json
// @Param bank formData model.Bank true "bank"
// @Param image formData file true "logo bank"
// @Security JWT
// @Success 201 {object} helper.Response{code=string,isError=bool,status=string,data=model.Bank}
// @Failure 400 {object} helper.Response{code=int,description=string,isError=bool}
// @Failure 500 {object} helper.Response{code=int,description=string,isError=bool}
// @Router /bank [post]
func (h *bankHandler) insert(c *gin.Context) {
	var request model.Bank
	if err := c.Bind(&request); err != nil {
		helper.GetResponse(err.Error(), 400, true).Send(c)
		return
	}

	err := h.validate.Struct(request)
	if err != nil {
		helper.GetResponse(err.Error(), http.StatusBadRequest, true).Send(c)
		return
	}

	url, err := helper.UploadImage(c)
	if err != nil {
		helper.GetResponse(err.Error(), 500, true).Send(c)
		return
	}

	request.LogoUrl = url

	h.insertBank.Execute(request).Send(c)
}
