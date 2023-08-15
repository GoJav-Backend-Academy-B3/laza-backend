package credit_card

import (
	"fmt"

	"github.com/phincon-backend/laza/domain/model"
	repo "github.com/phincon-backend/laza/domain/repositories"
	uc "github.com/phincon-backend/laza/domain/usecases/credit_card"
)

type getAllCreditCardUsecase struct {
	getAllCcRepo repo.GetAllAction[model.CreditCard]
}

func (uc *getAllCreditCardUsecase) Execute(userId uint64) (_result []model.CreditCard, err error) {
	rs, err := uc.getAllCcRepo.GetAll()
	if err != nil {
		return
	}

	for _, v := range rs {
		if v.UserId == userId {
			_result = append(_result, v)
			fmt.Println(v.UserId)
		}
	}
	return
}

func NewgetAllCreditCardUsecase(getAllCcRepo repo.GetAllAction[model.CreditCard]) uc.GetAllCreditCardUsecase {
	return &getAllCreditCardUsecase{
		getAllCcRepo: getAllCcRepo,
	}
}
