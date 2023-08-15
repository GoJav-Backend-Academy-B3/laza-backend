package credit_card

import (
	"github.com/phincon-backend/laza/domain/model"
	repo "github.com/phincon-backend/laza/domain/repositories"
	uc "github.com/phincon-backend/laza/domain/usecases/credit_card"
)

type getByIdCreditCardUsecase struct {
	getCcRepo repo.GetByIdAction[model.CreditCard]
}

func (h *getByIdCreditCardUsecase) Execute(ccId uint64) (_result model.CreditCard, err error) {
	_result, err = h.getCcRepo.GetById(ccId)
	return
}

func NewgetByIdCreditCardUsecase(getCcRepo repo.GetByIdAction[model.CreditCard]) uc.GetByIdCreditCardUsecase {
	return &getByIdCreditCardUsecase{
		getCcRepo: getCcRepo,
	}
}
