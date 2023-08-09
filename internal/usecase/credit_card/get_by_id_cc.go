package credit_card

import (
	"net/http"

	"github.com/phincon-backend/laza/domain/model"
	repo "github.com/phincon-backend/laza/domain/repositories"
	"github.com/phincon-backend/laza/domain/response"
	uc "github.com/phincon-backend/laza/domain/usecases/credit_card"
	"github.com/phincon-backend/laza/helper"
	"gorm.io/gorm"
)

type getByIdCreditCardUsecase struct {
	getCcRepo repo.GetByIdAction[model.CreditCard]
}

func (h *getByIdCreditCardUsecase) Execute(ccId uint64) *helper.Response {
	rp, err := h.getCcRepo.GetById(ccId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return helper.GetResponse(err.Error(), http.StatusNotFound, true)
		}
		return helper.GetResponse(err.Error(), http.StatusInternalServerError, true)
	}

	dt := response.CreditCardResponse{}.FillFromEntity(rp)
	return helper.GetResponse(dt, http.StatusOK, false)
}

func NewgetByIdCreditCardUsecase(getCcRepo repo.GetByIdAction[model.CreditCard]) uc.GetByIdCreditCardUsecase {
	return &getByIdCreditCardUsecase{
		getCcRepo: getCcRepo,
	}
}
