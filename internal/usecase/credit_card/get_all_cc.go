package credit_card

import (
	"net/http"

	"github.com/phincon-backend/laza/helper"

	"github.com/phincon-backend/laza/domain/model"
	repo "github.com/phincon-backend/laza/domain/repositories"
	"github.com/phincon-backend/laza/domain/response"
	uc "github.com/phincon-backend/laza/domain/usecases/credit_card"
)

type getAllCreditCardUsecase struct {
	getAllCcRepo repo.GetAllAction[model.CreditCard]
}

func (uc *getAllCreditCardUsecase) Execute(userId uint64) *helper.Response {
	rs, err := uc.getAllCcRepo.GetAll()
	if err != nil {
		return helper.GetResponse(err.Error(), http.StatusInternalServerError, true)
	}

	result := []response.CreditCardResponse{}
	for _, v := range rs {
		if v.UserId == userId {
			md := response.CreditCardResponse{}.FillFromEntity(v)
			result = append(result, md)
		}
	}
	return helper.GetResponse(result, http.StatusOK, false)
}

func NewgetAllCreditCardUsecase(getAllCcRepo repo.GetAllAction[model.CreditCard]) uc.GetAllCreditCardUsecase {
	return &getAllCreditCardUsecase{
		getAllCcRepo: getAllCcRepo,
	}
}
