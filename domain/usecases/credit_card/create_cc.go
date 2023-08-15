package credit_card

import (
	"github.com/phincon-backend/laza/domain/model"
	cc_request "github.com/phincon-backend/laza/domain/requests"
)

type AddCreditCardUsecase interface {
	Execute(userId uint64, md cc_request.CreditCardRequest) (_result model.CreditCard, statusCode int, err error)
}
