package credit_card

import (
	"github.com/phincon-backend/laza/domain/model"
)

type GetAllCreditCardUsecase interface {
	Execute(userId uint64) (_result []model.CreditCard, err error)
}
