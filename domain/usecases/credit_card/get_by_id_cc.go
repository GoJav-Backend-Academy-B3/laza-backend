package credit_card

import (
	"github.com/phincon-backend/laza/domain/model"
)

type GetByIdCreditCardUsecase interface {
	Execute(ccId uint64) (_result model.CreditCard, err error)
}
